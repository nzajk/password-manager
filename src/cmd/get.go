package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/nzajk/password-manager/src/db"
	"github.com/spf13/cobra"
)

// get a password from the database
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password from the database.",
	Run: func(cmd *cobra.Command, args []string) {
		// load the environment variables
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		loggedIn := os.Getenv("LOGGED_IN")
		// fmt.Println("LOGGED_IN:", loggedIn)

		// check if the user is logged in
		if loggedIn != "true" {
			log.Fatal("You must be logged in to get a password.")
		}

		if len(args) != 1 {
			log.Fatal("Usage: password-manager get <service>")
		}

		service := args[0]

		// connect to the database
		postgres, err := db.Connect()
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}

		// retrieve the password from the database
		row := postgres.QueryRow("SELECT password FROM passwords WHERE service=$1;", service)
		var password string
		err = row.Scan(&password)
		if err != nil {
			log.Fatalf("Error retrieving password: %v", err)
		}

		// todo: ensure that the same key is used for encryption and decryption (store the key safely)
		// password = crypto.Decrypt(password, crypto.GenerateKey(32))
		fmt.Println(password)

		// clean up
		postgres.Close()
	},
}
