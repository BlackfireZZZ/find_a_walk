package main

import (
	"database/sql"
	"github.com/gojuno/goose"
	"log"
	// Init DB drivers.
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	//  Migrations
	connectionString := "host=db port=5432 user=postgres password=PROD dbname=PROD sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Println(err)
	}
	if err := goose.Up(db, "/sql_migrations"); err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		log.Println(err)
	}
}
