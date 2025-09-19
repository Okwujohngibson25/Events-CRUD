package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/okwu-john/webapi/middlewares"
)

func Registerroutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group(("/"))
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvents)

	server.POST("/signup", createuser)
	server.POST("/login", loginuser)
}
