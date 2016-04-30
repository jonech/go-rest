package main

import (
	"log"
	"net/http"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// Change to your own local database setting
	DB_USER = "jchee"
	DB_PASSWORD = "1234"
	DB_NAME = "empty_dev"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getConnection() *sql.DB {

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", 
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func main() {

	router := NewRouter()

	// specify the port number here
	log.Fatal(http.ListenAndServe(":1234", router))
}

