package routes

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/okwu-john/webapi/models"
)

func registerforanevent(c *gin.Context) {
	userid := c.GetInt64("userid")
	eventid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Couldnt retrive/Parse Id", "ID": eventid})
		return

	}

	event, err := models.GetEventById(eventid)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			c.JSON(http.StatusNotFound, gin.H{"error": "event not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "database error"})
		}
		return
	}

	err = event.Register(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot register for Event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully registred for this event"})

}

func cancelregistration(c *gin.Context) {
	userid := c.GetInt64("userid")

	eventid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventById(eventid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	err = event.Cancelreg(userid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully cancelled registration for this event"})
}
