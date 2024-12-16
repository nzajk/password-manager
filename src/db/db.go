package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// connects to the database and returns the connection
func ConnectDB() (*sql.DB, error) {
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

	return db, nil
}

// queries the database and prints the results
func QueryDB(db *sql.DB, query string) {
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

// adds a new row to the database
func AddDB(db *sql.DB, query string) {
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
