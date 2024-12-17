package main

import (
	"fmt"

	"github.com/nzajk/password-manager/src/crypto"
	"github.com/nzajk/password-manager/src/db"
	"github.com/nzajk/password-manager/src/schemas"
)

func main() {
	plaintext := "testing to see if it works"
	key := crypto.GenerateKey(32)

	// encrypt the plaintext
	encrypted := crypto.Encrypt(plaintext, key)
	// fmt.Println("Encrypted:", encrypted)

	testEntry := schemas.Entry{
		ID:       2,
		Service:  "test",
		Username: "test",
		Password: encrypted,
	}

	postgres, err := db.Connect()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	db.AddRow(postgres, testEntry)
	db.Query(postgres, "SELECT * FROM passwords;")

	// clean up
	postgres.Close()
}
