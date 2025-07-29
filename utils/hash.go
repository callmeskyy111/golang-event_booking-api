package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// []bytes(str) = str->bytes
// string(bytes) = bytes -> str

// Hashing Utility fx #️⃣
func HashPassword(password string) (string,error) {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// Password Validating fx 🔐
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		fmt.Println("DEBUG: Password comparison failed:", err)
		return false
	}
	return true
}
