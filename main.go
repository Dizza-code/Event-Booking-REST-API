package main

import (
	"example.com/events-api/db"
	"example.com/events-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // call the DB function

	server := gin.Default()
	routes.RegisterRoutes(server) //passing the engine pointer to register route

	server.Run(":8080") //localhost: 8080
}
