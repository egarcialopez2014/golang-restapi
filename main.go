package main

import (
	"example.com/apirest/db"
	"example.com/apirest/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //  run the server on localhost:8080
}
