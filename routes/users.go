package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func SignUp(ctxt *gin.Context){
	var user models.User
	err := ctxt.ShouldBindJSON(&user)

if err != nil{
		ctxt.JSON(http.StatusBadRequest, gin.H{
		"message": "Could not parse request-data 🔴",
		"error": err.Error(),
	})
		return
}

	err =user.Save()

	if err!= nil{
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