package db

import (
	"database/sql"

	"goTestApi/config"

	_ "github.com/mattn/go-sqlite3"
)

// ConnectDB connects to desired database
func ConnectDB(config *config.Config) *sql.DB {
	// connect to DB
	db, err := sql.Open(
		"sqlite3", config.DBNAME,
	)
	checkErr(err)

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}

// checkErr checks for errors and panics if found
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
