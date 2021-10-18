package main

import (
	"fmt"

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
	d := db.ConnectDB(config.New())
	fmt.Println(d)
}
