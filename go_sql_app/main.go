package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Setup database connection
	var err error

	connect_string := fmt.Sprintf("%s:%s@tcp(localhost:3306)/gitops", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"))
	db, err = sql.Open("mysql", connect_string)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup HTTP server
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Run a simple query
	rows, err := db.Query("SELECT * FROM people;")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Process query result
	for rows.Next() {

		var firstname string
		var lastname string
		var address string
		var age int

		if err := rows.Scan(&firstname, &lastname, &age, &address); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "firstname, lastname, age, address\n")
		fmt.Fprintf(w, "%s, %s, %d, %s\n", firstname, lastname, age, address)
	}

	// Check for errors from iterating over result
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
