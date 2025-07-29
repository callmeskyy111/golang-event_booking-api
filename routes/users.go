package routes

import (
	"fmt"
	"net/http"

	//"strings"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func SignUp(ctxt *gin.Context){
	var user models.User
	err := ctxt.ShouldBindJSON(&user)

if err != nil{
	 	fmt.Println("DEBUG - Error during binding:", err)
		ctxt.JSON(http.StatusBadRequest, gin.H{
		"message": "Could not parse request-data 🔴",
		"error": err.Error(),
	})
		return
}

	err =user.Save()

	if err!= nil{
		fmt.Println("DEBUG - Error during binding:", err)
		ctxt.JSON(http.StatusInternalServerError, gin.H{"message":"Could not save USER.. Try again later 🔴", "success":"false"})
		return
	}
	ctxt.JSON(http.StatusCreated, gin.H{
	"message": "User created successfully✅",
	"success": "true",
	"user": gin.H{
		"userId":    user.Id,
		"email": user.Email,
	},
})

}



func Login(ctxt *gin.Context){
	var user models.User
	err := ctxt.ShouldBindJSON(&user)

	if err != nil{
		ctxt.JSON(http.StatusBadRequest, gin.H{
		"message": "Could not parse request-data 🔴",
		"error": err.Error(),
	})
		return
}

err = user.ValidateCredentials()

if err != nil{
	ctxt.JSON(http.StatusUnauthorized,gin.H{"message":"User-Authentication Error 🔴", "success":"false"})
	return
}

token,err:= utils.GnerateToken(user.Email, user.Id)

if err != nil{
	ctxt.JSON(http.StatusInternalServerError,gin.H{"message":"User-Authentication Error 🔴", "success":"false"})
	return
}

ctxt.JSON(http.StatusOK, gin.H{"message":"Login Successful ✅", "token":token})
}

