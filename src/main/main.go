package main

func main() {
	plaintext := "Hello"
	key := crypto.generateKey(32)

	encrypted := crypto.encrypt(plaintext, key)
	println(encrypted)
	decrypted := crypto.decrypt(encrypted, key)
	println(decrypted)
}
