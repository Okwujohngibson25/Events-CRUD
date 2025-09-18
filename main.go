package main

import (
	"github.com/gin-gonic/gin"
	"github.com/okwu-john/webapi/db"
	"github.com/okwu-john/webapi/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.Registerroutes(server)

	server.Run(":8080")
}
