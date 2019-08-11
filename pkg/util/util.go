package util

import "golang.org/x/crypto/bcrypt"

// SaltHashGenerate encrypts user passwords
func SaltHashGenerate(password string) (string, error) {
	hex := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(hex, 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// SaltHashCompare compares passwords for consistency
func SaltHashCompare(digest []byte, password string) bool {
	hex := []byte(password)

	if err := bcrypt.CompareHashAndPassword(digest, hex); err == nil {
		return true
	}

	return false
}
