package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okwu-john/webapi/models"
	"github.com/okwu-john/webapi/utils"
)

func createuser(c *gin.Context) {

	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not parse in value"})
	}

	err = newUser.Createuser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not insert into users DB", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user created"})

}

func loginuser(c *gin.Context) {
	var newUser models.User
	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "could not parse incoming data"})
		return
	}

	err = newUser.Validateuserlogin()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password validating failed"})
		return
	}

	token, err := utils.GenerateToken(newUser.Email, newUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cant generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged in!",
		"token":   token,
	})

}
