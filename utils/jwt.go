package utils

import (
	"errors"
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


func VerifyToken(token string)error{
	parsedToken,err:=jwt.Parse(token, func(token *jwt.Token)(any,error){
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil, errors.New("err! 🔴 Unexpected sign-in method")
		}
		return []byte(secretKey),nil
	})
	if err!=nil{
		return  errors.New("err! 🔴 Could not parse token")
	}
	tokenIsValid := parsedToken.Valid

	if !tokenIsValid{
		return errors.New("error! 🔴 Invalid Token")
	}

	//! Extracting fields (Optnl.) ✅
	//claims,ok:= parsedToken.Claims.(jwt.MapClaims)

	// if !ok{
	// 	return errors.New("error! 🔴 Invalid Token Claims")
	// }

	// email:=claims["email"].(string)
	// userId:=claims["userId"].(int64)

	return nil
}