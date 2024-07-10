package main

import (
	"database/sql"
	"github.com/pressly/goose/v3"
	"log"
	// Init DB drivers.
	_ "github.com/lib/pq"
)

func main() {
	connectionString := "host=db port=5432 user=postgres password=PROD dbname=PROD sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}
}
