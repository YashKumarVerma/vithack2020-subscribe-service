package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

func main() {
	// connecting application to database
	log.Println("Attempting to connect to SQL Database")

	db, _ = sql.Open("mysql", "yash:yashpassword@/vithack2020")

	if db == nil {
		log.Fatalln("Error connecting to database")
	} else {
		log.Println("Database Connected")
	}

	defer db.Close()
	defer log.Println("Database Disconnected")

	// mounting routes
	route := mux.NewRouter()
	route.HandleFunc("/new", addListener).Methods("POST")

	// starting server
	log.Println("Starting server on http://127.0.0.1:8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}
