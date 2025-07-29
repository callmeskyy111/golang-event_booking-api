package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterForEvent(ctxt *gin.Context) {
	userId := ctxt.GetInt64("userId")
	eventId,err := strconv.ParseInt(ctxt.Param("id"),10,64)

	
	if err!=nil {
		ctxt.JSON(http.StatusBadRequest, gin.H{"message":"Could not parse the Event-ID.. Try again later 🔴", "success":"false"})
		return
	}

	event,err:=models.GetEventById(eventId)
	if err!=nil{
		ctxt.JSON(http.StatusInternalServerError, gin.H{"message":"Could Not Fetch Event 🔴","success":"false"})
		return
	}

	err = event.Register(userId)

	if err!=nil{
			ctxt.JSON(http.StatusInternalServerError, gin.H{"message":"Could Not REGISTER User For The Event 🔴","success":"false"})
		return
	}
	ctxt.JSON(http.StatusCreated, gin.H{"message":"Registered ✅"})
}

func CancelRegistration(ctxt *gin.Context) {
userId := ctxt.GetInt64("userId")
eventId, err := strconv.ParseInt(ctxt.Param("id"),10,64)
var event models.Event
event.ID = eventId

err = event.Cancel(userId)

if err!=nil{
			ctxt.JSON(http.StatusInternalServerError, gin.H{"message":"Could Not Cancel Registration 🔴","success":"false"})
		return
	}
	ctxt.JSON(http.StatusOK, gin.H{"message":"Registration Cancelled ✅"})


}