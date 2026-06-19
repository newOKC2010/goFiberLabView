package database

import (
	"database/sql"
	"log"

	loadenv "view_lab/src/loadenv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() *sql.DB {
	if DB == nil {
		loadenv.LoadDBconnec()

		dbURL, err := validateAndGetDBURL()
		if err != nil {
			log.Fatal(err)
		}

		DB, err = sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatal(err)
		}

		if err = DB.Ping(); err != nil {
			log.Fatal(err)
		}

		log.Println("connect database success")
		HandleModelCreation()
	}
	return DB
}
