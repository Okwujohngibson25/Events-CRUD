package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/okwu-john/webapi/models"
)

func getEvents(c *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Couldnt fetch rows", "error": err.Error()})

	}
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized access", "err": "No token in request"})
		return
	}

	var event models.Event

	event.ID = 1
	event.UserID = 1

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Binding incomming request to my struct var didnt work", "error": err.Error()})
		return
	}

	err = event.SaveEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "couldnt create event", "error": err.Error()})

	}

	c.JSON(http.StatusOK, gin.H{"message": "Created event", "event": event})
}

func getEvent(c *gin.Context) {

	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not pass result ID", "error": err.Error()})
		return
	}

	result, err := models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Couldn't find event using ID", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gotten ID", "event": result})

}
func updateEvent(c *gin.Context) {
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "could not pass event to ID"})
		return
	}
	_, err = models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "could find event", "error": err})
		return
	}

	var updatedevent models.Event
	err = c.ShouldBindJSON(&updatedevent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Binding incomming request to my struct var didnt work", "error": err.Error()})
		return
	}

	updatedevent.ID = eventId

	err = updatedevent.UpdateEvent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Cant update rows", "error": err.Error()})
		return

	}

	c.JSON(http.StatusOK, gin.H{"messgae": "Event Updated successfully", "ID": eventId})
}

func deleteEvents(c *gin.Context) {
	eventID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt retrive/Parse Id", "ID": eventID})
		return

	}

	_, err = models.GetEventById(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt find an Event with Passed in ID"})
		return
	}

	err = models.DeleteEvents(eventID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error in deleting event"})

	}

	c.JSON(http.StatusOK, gin.H{"message": "event was deleted"})
}
