package main

import (
	"fmt"

	"example.com/events-api/db"
	"example.com/events-api/routes"
	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB() // call the DB function

	server := gin.Default()
	routes.RegisterRoutes(server) //passing the engine pointer to register route

	server.Run(":8080") //localhost: 8080

	// The plain text password you want to update
	plainPassword := "testing"

	// Generate bcrypt hash
	hashedPassword, err := utils.HashPassword(plainPassword)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return
	}

	// Print the hash to use in the SQL query
	fmt.Println("Hashed Password:", hashedPassword)
}
