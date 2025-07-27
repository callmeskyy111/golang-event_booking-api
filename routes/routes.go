package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
server.GET("/events", GetEvents) 
server.GET("/events/:id", GetSingleEvent) 
server.POST("/events", CreateEvent)
server.PUT("/events/:id",UpdateEvent)
}