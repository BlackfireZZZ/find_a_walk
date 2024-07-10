package main

import (
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
	"log"
	"time"

	// Init DB drivers.
	_ "github.com/lib/pq"
)

var embedMigrations embed.FS

func main() {
	connectionString := "host=db port=5432 user=postgres password=PROD dbname=PROD sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	time.Sleep(40 * time.Second)
	if err := goose.Up(db, "."); err != nil {
		log.Fatal(err)
	}
}
