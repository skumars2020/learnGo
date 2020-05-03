package model

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:1234@(tcp:localhost:3306)")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
