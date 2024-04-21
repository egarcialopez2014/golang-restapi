package main

import (
	"net/http"

	"example.com/apirest/db"
	"example.com/apirest/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) // the routes and the handlers
	server.POST("/events", createEvent)

	server.Run(":8080") //  run the server on localhost:8080
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // analog to scan but for gin framework

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request."})
		return
	}

	event.ID = 1
	event.UserID = 4

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
