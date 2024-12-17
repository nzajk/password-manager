package cmd

import (
	"fmt"
	"log"

	"github.com/nzajk/password-manager/src/crypto"
	"github.com/nzajk/password-manager/src/db"
	"github.com/spf13/cobra"
)

// get a password from the database
var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password from the database.",
	Run: func(cmd *cobra.Command, args []string) {
		if !loggedIn {
			log.Fatal("You must be logged in to get a password.")
		}

		if len(args) != 1 {
			log.Fatal("Usage: password-manager get <service>")
		}

		service := args[0]

		postgres, err := db.Connect()
		if err != nil {
			fmt.Println("Error connecting to the database:", err)
		}

		row := postgres.QueryRow("SELECT password FROM passwords WHERE service=$1;", service)
		var password string
		err = row.Scan(&password)
		if err != nil {
			fmt.Println("Error retrieving password:", err)
		}

		password = crypto.Decrypt(password, crypto.GenerateKey(32))
		fmt.Println(password)

		postgres.Close()
	},
}
