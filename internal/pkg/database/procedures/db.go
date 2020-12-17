package procedures

import (
	"github.com/ReolinkCameraAPI/noctilucago/internal/pkg/database/models"
	_ "gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
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

	if os.Getenv("NOCTI_DB") == "" {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open(os.Getenv("NOCTI_DB")), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	// Get the generic DB interface
	sqlDb, err := db.DB()

	if err != nil {
		return nil, err
	}

	sqlDb.SetMaxIdleConns(maxIdleConns)
	sqlDb.SetConnMaxLifetime(maxConnLifetime)

	return &DB{db}, nil
}

func (db *DB) Migrate() error {
	var tables []interface{}

	tables = append(tables, models.Camera{})
	tables = append(tables, models.CameraModel{})
	tables = append(tables, models.CameraLocation{})

	tables = append(tables, models.Proxy{})

	if err := db.Debug().AutoMigrate(tables...); err != nil {
		return err
	}

	return nil
}
