package utils

import "golang.org/x/crypto/bcrypt"

// []bytes(str) = str->bytes
// string(bytes) = bytes -> str

// Hashing Utility fx #️⃣

func HashPassword(password string) (string,error) {
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}