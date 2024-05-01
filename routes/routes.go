package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// dont need caps because part of the same package
	server.GET("/events", getAllEvents) // the routes and the handlers
	server.GET("/events/:id", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
}
