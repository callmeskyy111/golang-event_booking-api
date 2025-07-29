package middlewares

import (
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctxt *gin.Context) {
token:= ctxt.Request.Header.Get("Authorization")
	if token==""{
		ctxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Authorization-Error 🔴", "success":"false"})
		return
	}

	userId,err:=utils.VerifyToken(token)

	if err!=nil{
		ctxt.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message":"Not Authorized 🔴","success":"false"})
		return
	}	
	ctxt.Set("userId",userId)
	ctxt.Next()
}