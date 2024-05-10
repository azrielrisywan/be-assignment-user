package config

import (
	"log"
	"github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func SetupDatabase() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=Tasik123 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	 // Test the connection to the database
	if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Database Successfully Connected")
    }
	return db
}