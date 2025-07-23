package main

import (
	"net/http"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
server:=gin.Default()
server.GET("/events", getEvents) 
server.POST("/events", createEvent)
server.Run(":8080")
}

func getEvents(context *gin.Context){
	events:=models.GetAllEvents()
	context.JSON(http.StatusOK,events)
}

func createEvent(ctxt *gin.Context){
	var event models.Event
	err :=ctxt.ShouldBindJSON(&event)

	if err != nil{
		ctxt.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse request-data 🔴", "success":"false"})
		return
	}
	
	event.ID = 1
	event.UserId = 1

	event.Save()

	ctxt.JSON(http.StatusCreated,gin.H{"message":"Event created✅", "success":"true", "event":event})
}