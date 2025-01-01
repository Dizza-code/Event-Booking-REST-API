package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //setting up a dynamic path handler
	server.POST("/events", createEvent)
	server.PUT("events/:id", updateEvent)
	server.DELETE("events/:id", deleteEvent)
}
