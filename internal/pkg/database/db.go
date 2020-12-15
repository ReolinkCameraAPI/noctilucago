package database

import (
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

func Migrate() error {
	var err error
	var db *gorm.DB
	// var dbConstraints *gorm.DB

	var tables []interface{}

	maxIdleConns := 0
	maxConnLifetime := time.Minute * 2

	if os.Getenv("NOCTI_DB") == "" {
		db, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	} else {
		db, err = gorm.Open(sqlite.Open(os.Getenv("NOCTI_DB")), &gorm.Config{})
	}

	if err != nil {
		return err
	}

	// Get the generic DB interface
	sqlDb, err := db.DB()

	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(maxIdleConns)
	sqlDb.SetConnMaxLifetime(maxConnLifetime)

	if err := db.Debug().AutoMigrate(tables...); err != nil {
		return err
	}

	return nil
}
