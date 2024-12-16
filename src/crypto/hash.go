package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
