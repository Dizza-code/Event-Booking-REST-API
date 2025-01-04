package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/events-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("DEBUG: Parse error: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		fmt.Printf("DEBUG: GetEventByID error: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		fmt.Printf("DEBUG: Register error: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered "})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		fmt.Printf("DEBUG: Parse error: %v\n", err)
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		fmt.Printf("DEBUG: Register error: %v\n", err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel event for user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled"})
}

func getRegistrations(context *gin.Context) {
	registrations, err := models.GetAllRegistrations()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch registrations"})
		return
	}
	context.JSON(http.StatusOK, registrations)
}

// function for a single registration
func getRegistration(context *gin.Context) {
	registrationId, err := strconv.ParseInt(context.Param("id"), 10, 64) // to get the parameter value
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch registration"})
		return
	}
	registration, err := models.GetRegistrationByID(registrationId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse registration id"})
		return
	}
	context.JSON(http.StatusOK, registration)

}
