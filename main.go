package main

import (
	"goTestApi/config"
	"goTestApi/db"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	// initialize DB
	d, err := db.NewDB(config.New())
	if err != nil {
		panic(err)
	}
	defer d.Close()

}
