package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func queryDB(db *sql.DB, query string) {
	rows, err := db.Query(query)
	if err != nil {
		print("Error with query:", err)
		return
	}

	// iterate through the rows and print the data
	for rows.Next() {
		// table structure: id, service, username, password
		var id int
		var service string
		var username string
		var password string

		err := rows.Scan(&id, &service, &username, &password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id: %d, service: %s, username: %s, password: %s\n", id, service, username, password)
	}
}

func main() {
	// load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get connection details from environment variables
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")

	// construct connection string
	connectionStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}

	queryDB(db, "SELECT * FROM passwords;")

	db.Close()
}
