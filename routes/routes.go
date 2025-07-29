package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
//All CRUD	
server.GET("/events", GetEvents) 
server.GET("/events/:id", GetSingleEvent) 

// Grouping for the middleware
authenticated:=server.Group("/")
authenticated.Use(middlewares.Authenticate)
authenticated.POST("/events",CreateEvent)
authenticated.PUT("/events/:id",UpdateEvent)
authenticated.DELETE("/events/:id",DeleteEvent)
authenticated.POST("/events/:id/register",RegisterForEvent)
authenticated.DELETE("/events/:id/register",CancelRegistration)

server.POST("/signup", SignUp)
server.POST("/login", Login)
}
