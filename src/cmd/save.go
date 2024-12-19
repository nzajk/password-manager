package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nzajk/password-manager/src/crypto"
	"github.com/nzajk/password-manager/src/db"
	"github.com/spf13/cobra"
)

// saves a password to the database
var SaveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a password to the database.",
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
			log.Fatal("You must be logged in to save a password.")
		}

		if len(args) != 3 {
			log.Fatal("Usage: password-manager save <service> <username> <password>")
		}

		service := args[0]
		username := args[1]
		password := args[2]

		password = crypto.Encrypt(password, crypto.GenerateKey(32))

		postgres, err := db.Connect()
		if err != nil {
			fmt.Println("Error connecting to the database:", err)
		}

		_, err = postgres.Exec("INSERT INTO passwords (service, username, password) VALUES ($1, $2, $3);", service, username, password)
		if err != nil {
			fmt.Println("Error inserting into the database:", err)
		}

		postgres.Close()
	},
}
