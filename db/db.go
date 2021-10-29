package db

import (
	"time"

	"github.com/mdmims/go-echo-crud/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// NewDB connects to desired database
func NewDB(config *config.Config) (*sqlx.DB, error) {
	// connect to DB
	db, err := sqlx.Open("sqlite3", config.DBNAME)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	// set reasonable connection limits
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
