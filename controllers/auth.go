package controllers

import (
	"net/http"

	"github.com/garylow2001/GossipGo-Backend/initializers"
	"github.com/garylow2001/GossipGo-Backend/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(context *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if context.Bind(&body) != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to hash password"})
		return
	}

	user := models.User{Username: body.Username}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create user, please use another username"})
		return
	}

	auth := models.Auth{UserID: user.ID, Password: string(hash)}

	result = initializers.DB.Create(&auth)

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create auth for user"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "User created"})
}
