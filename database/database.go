package database

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB
var host = os.Getenv("POSTGRES_HOST")
var port = os.Getenv("POSTGRES_PORT")
var postgresUser = os.Getenv("POSTGRES_USER")
var password = os.Getenv("POSTGRES_PASSWORD")
var dbname = os.Getenv("POSTGRES_DATABASE")

func Connection() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, postgresUser, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("error em login.sql.Open ", err)
	}

	// close database

	// check db
	err = db.Ping()
	if err != nil {
		fmt.Println("error em login.db.Ping ", err)
	}

	fmt.Println("Connected!")
	return db
}
