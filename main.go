package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"example.com/rest-api/db"
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main(){
	db.InitDB()
server:=gin.Default()
server.GET("/events", getEvents) 
server.GET("/events/:id", getSingleEvent) 
server.POST("/events", createEvent)
server.Run(":8080")
}

func getSingleEvent(ctxt *gin.Context){
	eventId,err:=strconv.ParseInt(ctxt.Param("id"), 10, 64) 

	if err!=nil {
		ctxt.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse the Event-ID.. Try again later 🔴", "success":"false"})
		return
	}
	event,err:=models.GetEventById(eventId)
	// if err!=nil {
	// 	ctxt.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch the Event-ID.. Try again later 🔴", "success":"false"})
	// 	return
	// }
	if err != nil {
	if errors.Is(err, sql.ErrNoRows) {
		ctxt.JSON(http.StatusNotFound, gin.H{"message": "Event not found 🔍"})
		return
	}
	ctxt.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the Event 🔴"})
	return
}
	ctxt.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context){
	events,err:=models.GetAllEvents()
	if err!=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Could not fetch events.. Try again later 🔴", "success":"false"})
		return
	}
	context.JSON(http.StatusOK,events)
}

func createEvent(ctxt *gin.Context){
	var event models.Event
	err := ctxt.ShouldBindJSON(&event)
if err != nil {
	fmt.Println("🛑 Bind error:", err)
	ctxt.JSON(http.StatusBadRequest, gin.H{
		"message": "Could not parse request-data 🔴",
		"error": err.Error(),
	})
	return
}
	
	//event.ID = 1
	event.UserId = 1

	err = event.Save()

	if err!= nil{
		ctxt.JSON(http.StatusInternalServerError, gin.H{"message":"Could not create event.. Try again later 🔴", "success":"false"})
		return
	}

	ctxt.JSON(http.StatusCreated,gin.H{"message":"Event created✅", "success":"true", "event":event})
}
