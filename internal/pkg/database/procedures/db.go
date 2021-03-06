package procedures

import (
	"fmt"
	"github.com/ReolinkCameraAPI/noctilucago/config"
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/url"
	"strings"
	"time"
)

type DB struct {
	*gorm.DB
}

func NewDatabase() (*DB, error) {
	var db *gorm.DB
	var err error

	maxIdleConns := 2
	maxConnLifetime := time.Hour * 1

	var dsn string
	parsed, err := url.Parse(config.NlConfig.DSN)

	if err != nil {
		return nil, err
	}

	username := parsed.User.Username()
	password, _ := parsed.User.Password()
	host := parsed.Host
	var port string

	if strings.Contains(host, ":") {
		port = strings.Split(host, ":")[1]
		host = strings.Split(host, ":")[0]
	}

	scheme := parsed.Scheme
	database := strings.Trim(parsed.Path, "/")
	queries := parsed.Query()

	//var sslmode string
	//var sslrootcert string
	//var sslcert string
	//var sslkey string

	var extra string

	switch scheme {
	case "postgresql":
		if port == "" {
			port = "5432"
		}

		host = fmt.Sprintf("host=%s", host)
		port = fmt.Sprintf("port=%s", port)

		dsn = fmt.Sprintf("%s %s", host, port)

		if database != "" {
			dsn = fmt.Sprintf("%s dbname=%s", dsn, database)
		}

		if username != "" {
			dsn = fmt.Sprintf("%s user=%s", dsn, username)
		}

		if password != "" {
			dsn = fmt.Sprintf("%s password=%s", dsn, password)
		}

		for key, val := range queries {
			extra = fmt.Sprintf("%s %s=%s", extra, key, val[0])
		}

		if extra != "" {
			dsn = fmt.Sprintf("%s %s", dsn, extra)
		}

		/*if sslmode = queries.Get("sslmode"); sslmode == "" {
			sslmode = "disabled"
		}

		ssl = fmt.Sprintf("sslmode=%s", sslmode)

		if sslrootcert = queries.Get("sslrootcert"); sslrootcert != "" {
			ssl = fmt.Sprintf("%s sslrootcert=%s", ssl, sslrootcert)
		}

		if sslcert = queries.Get("sslcert"); sslcert != "" {
			ssl = fmt.Sprintf("%s sslcert=%s", ssl, sslcert)
		}

		if sslkey = queries.Get("sslkey"); sslkey != "" {
			ssl = fmt.Sprintf("%s sslkey=%s", ssl, sslkey)
		}*/

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		break
	case "mysql":
		// TODO: add mysql support
		dsn = "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		panic("mysql is unsupported")
		break
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(host), &gorm.Config{})
		break
	default:
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
		break
	}

	if err != nil {
		return nil, err
	}
	dbWrapper := &DB{db}

	if scheme == "sqlite" {
		err = dbWrapper.Migrate()
		if err != nil {
			return nil, err
		}
	}

	// Get the generic DB interface
	sqlDb, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(maxIdleConns)
	sqlDb.SetConnMaxLifetime(maxConnLifetime)

	return dbWrapper, nil
}

func (db *DB) Migrate() error {
	var tables []interface{}

	tables = append(tables, models.User{})

	tables = append(tables, models.Camera{})
	tables = append(tables, models.CameraModel{})
	tables = append(tables, models.CameraAuth{})

	tables = append(tables, models.Proxy{})

	tables = append(tables, models.Location{})

	if err := db.Debug().AutoMigrate(tables...); err != nil {
		return err
	}

	return nil
}
