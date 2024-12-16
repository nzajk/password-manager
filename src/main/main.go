package main

import (
	"fmt"

	"github.com/nzajk/password-manager/src/crypto"
)

func main() {
	plaintext := "testing to see if it works"
	key := crypto.GenerateKey(32)

	// encrypt the plaintext
	encrypted := crypto.Encrypt(plaintext, key)
	fmt.Println("Encrypted:", encrypted)

	// decrypt the ciphertext
	decrypted := crypto.Decrypt(encrypted, key)
	fmt.Println("Decrypted:", decrypted)
}
