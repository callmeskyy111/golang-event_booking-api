package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecre1234"

func GnerateToken(email string, userId int64) (string,error){
	token:= jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp":time.Now().Add(time.Hour*2).Unix(),
		// Don't put password, for security reasons.
	})

	return token.SignedString([]byte(secretKey))
}