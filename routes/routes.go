package routes

import (
	"example.com/events-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //setting up a dynamic path handler
	server.GET("/registrations", getRegistrations)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate) //to protect the group
	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)
	authenticated.POST("events/:id/register", registerForEvent)
	authenticated.DELETE("events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
