package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) string {
	password := []byte(p)

	hashed, err := bcrypt.GenerateFromPassword(password, 8)
	if err != nil {
		panic(err)
	}
	return string(hashed)
}

func ComparePassword(p string, s string) bool {
	password := []byte(p)
	hashed := []byte(s)

	err := bcrypt.CompareHashAndPassword(hashed, password)
	return err == nil
}
