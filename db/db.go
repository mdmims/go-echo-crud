package db

import (
	"database/sql"

	"goTestApi/config"
)

// NewDB connects to desired database
func NewDB(config *config.Config) (*sql.DB, error) {
	// connect to DB
	sqliteDb, err := sql.Open("sqlite3", config.DBNAME)
	if err != nil {
		return nil, err
	}

	if err = sqliteDb.Ping(); err != nil {
		panic(err)
	}

	return sqliteDb, nil
}
