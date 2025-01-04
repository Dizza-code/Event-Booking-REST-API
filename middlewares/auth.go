package middlewares

import (
	"net/http"

	"example.com/events-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}
	userId, err := utils.VerifyToken(token) //verifying if the token is a valid token
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	context.Set("userId", userId) //this is a method that allows you add data to the context value, the datat is then attached to the context and can be used anywhere the context is available
	context.Next()                // this ensures that the next event handler in line will execute correctly
}
