package routes

import (
	"net/http"
	"strconv"

	"example.com/apirest/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func getAllEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events"})
	}

	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
	}

	context.JSON(http.StatusOK, event)

}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // analog to scan but for gin framework

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request."})
		return
	}

	event.UserID = 4

	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {

	// first parse the event ID
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event Id"})
		return
	}

	// then fetch that event
	_, err = models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent) // analog to scan but for gin framework

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request."})
		return
	}

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.Render(http.StatusOK, render.JSON{Data: any(gin.H{"message": "event has been updated"})})

}
