package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/okwu-john/webapi/utils"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized access", "err": "No token in request"})
		return
	}

	userid, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized access", "err": err.Error(), "token": token})
		return
	}
	c.Set("userid", userid)
	c.Next()
}
