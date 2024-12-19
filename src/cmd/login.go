package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/nzajk/password-manager/src/crypto"
	"github.com/nzajk/password-manager/src/db"
	"github.com/spf13/cobra"
)

// user must give a correct master password to login
var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the system.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Usage: password-manager login <username> <password>")
		}

		username := args[0]
		password := args[1]

		postgres, err := db.Connect()
		if err != nil {
			fmt.Println("Error connecting to the database:", err)
			return
		}

		// get the master password from the database
		row := postgres.QueryRow("SELECT password FROM master WHERE username=$1;", username)
		var masterPass string
		err = row.Scan(&masterPass)
		if err != nil {
			fmt.Println("Error retrieving password:", err)
			return
		}

		// hash the given password to compare with the master password
		if crypto.Hash(password) == masterPass {
			fmt.Println("Successfully logged in")
			os.Setenv("LOGGED_IN", "true")
		} else {
			fmt.Println("Incorrect password")
		}

		// clean up
		postgres.Close()
	},
}
