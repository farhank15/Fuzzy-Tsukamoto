package utils

import "golang.org/x/crypto/bcrypt"

const bcryptCost = 10 // Default cost factor is 10, you can lower it to 8 or 9 for faster hashing

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
