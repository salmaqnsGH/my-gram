package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	salt := 10
	arrByte := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(arrByte, salt)

	return string(hash), err
}
