package utils

import "golang.org/x/crypto/bcrypt"

//TODO: use argon2 to hash passwords to ensure performance

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hash), err
}

func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	isValid := err == nil

	return isValid
}
