package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"example.com/events-api/models"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch events. Try again later"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// function for getting a single event
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // to get your path paramenter value
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch events. Try again later"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "could not parse event ID"})
		return
	}
	context.JSON(http.StatusOK, event) // to show the success of the the ID that was fetched
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Could not parse request data"})
		return
	}
	event.ID = 1
	event.UserID = 1

	err = event.Save() // saving the event could also fail so we can handle a potentioal err by storing the response in the err variable
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create events. Try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}
func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // to get your path paramenter value
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch events. Try again later"})
		return
	}

	_, err = models.GetEventByID(eventId) //fecth the event that has the event id

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "could not fetch event"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBind(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "could not update event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "request updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // to get your path paramenter value
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not fetch events. Try again later"})
		return
	}
	event, err := models.GetEventByID(eventId) //fecth the event that has the event id

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "could not fetch event"})
		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
