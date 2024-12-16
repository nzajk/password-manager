package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// generate a key of proper length for AES
func generateKey(size int) []byte {
	key := make([]byte, size)
	rand.Read(key)
	return key
}

// encrypt function using AES-256
func encrypt(data string, key []byte) string {
	// ensure the key length is correct
	if len(key) != 32 {
		log.Fatal("Invalid key length. AES requires a 32-byte key.")
		return "Invalid key length."
	}

	plaintext := []byte(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatal(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		log.Fatal(err)
	}

	// encrypt the data using the nonce and GCM mode
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// convert the ciphertext to a hexadecimal string
	encrypted := hex.EncodeToString(ciphertext)

	return encrypted
}

// decrypt function using AES-256
func decrypt(ciphertext string, key []byte) string {
	// ensure the key length is correct (32 bytes for AES-256)
	if len(key) != 32 {
		log.Fatal("Invalid key length. AES requires a 32-byte key.")
		return "Invalid key length."
	}

	// decode the hexadecimal ciphertext
	decoded, err := hex.DecodeString(ciphertext)
	if err != nil {
		fmt.Println("error decoding hex:", err)
		return ""
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("error creating cipher block:", err)
		return ""
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("error creating GCM cipher mode:", err)
		return ""
	}

	// extract the nonce and the ciphertext
	nonce := decoded[:gcm.NonceSize()]
	ciphertextData := decoded[gcm.NonceSize():]

	// decrypt the data using the nonce and GCM mode
	plaintext, err := gcm.Open(nil, nonce, ciphertextData, nil)
	if err != nil {
		fmt.Println("error decrypting data:", err)
		return ""
	}

	return string(plaintext)
}
