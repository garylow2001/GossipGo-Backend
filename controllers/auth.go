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

	user := models.User{Username: body.Username, Auth: models.Auth{Password: string(hash)}}

	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Failed to create user"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "User created"})
}
