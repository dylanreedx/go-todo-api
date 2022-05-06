package utils

import "golang.org/x/crypto/bcrypt"

// It takes a password and a hash, and returns true if the password matches the hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
