package routes

import (
	"fmt"
	"net/http"

	"example.com/events-api/db"
	"example.com/events-api/models"
	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Could not parse request data"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func login(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Could not parse request data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		fmt.Println("Error parsing request:", err)
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not authenticate user"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})
}

func UpdatePassword(context *gin.Context) {
	var user struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	// Hash the new password
	hashedPassword, err := utils.HashPassword(user.NewPassword)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error hashing password"})
		return
	}

	// Update the password in the database
	err = db.UpdateUserPassword(user.Email, hashedPassword)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update password"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
