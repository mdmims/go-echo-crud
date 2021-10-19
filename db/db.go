package db

import (
	"database/sql"

	"github.com/mdmims/go-echo-crud/config"

	_ "github.com/mattn/go-sqlite3"
)

// NewDB connects to desired database
func NewDB(config *config.Config) (*sql.DB, error) {
	// connect to DB
	sqliteDb, err := sql.Open("sqlite3", config.DBNAME)
	if err != nil {
		return nil, err
	}

	if err = sqliteDb.Ping(); err != nil {
		return nil, err
	}

	return sqliteDb, nil
}
