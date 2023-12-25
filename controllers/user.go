package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/garylow2001/GossipGo-Backend/models"
	"github.com/gin-gonic/gin"
)

var Users []models.User

// User handlers
func GetUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Users)
}

func CreateUser(context *gin.Context) {
	var newUser models.User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	Users = append(Users, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

func GetUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"}) //TODO: abstract out invalid integer error message
		return
	}

	user, err := getUserByID(uint(id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"}) //TODO: abstract out error message
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func getUserByID(id uint) (*models.User, error) {
	for i, u := range Users {
		if u.ID == id {
			return &Users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func UpdateUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	user, err := getUserByID(uint(id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	var updatedUser models.User

	if err := context.BindJSON(&updatedUser); err != nil {
		return
	}

	user.Username = updatedUser.Username
	user.Password = updatedUser.Password

	context.IndentedJSON(http.StatusCreated, updatedUser)
}

func DeleteUser(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	user, err := getUserByID(uint(id))

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	Users = removeUser(Users, user)

	context.IndentedJSON(http.StatusCreated, Users)
}

func removeUser(users []models.User, user *models.User) []models.User {
	for i, u := range users {
		if u.ID == user.ID {
			return append(users[:i], users[i+1:]...)
		}
	}

	return users
}
