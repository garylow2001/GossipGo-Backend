package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/garylow2001/GossipGo-Backend/initializers"
	"github.com/garylow2001/GossipGo-Backend/models"
	"github.com/gin-gonic/gin"
)

var Users []models.User
var User models.User

// User handlers

/*
* GetUsers retrieves all users from the database and returns them as a JSON response.
* It does not show the posts and comments of the users.
  - @param context The context of the request.
*/
func GetUsers(context *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)

	if result.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No user found"})
		return
	}

	context.IndentedJSON(http.StatusOK, users)
}

/*
* CreateUser creates a new user and adds it to the database.
  - @param context The context of the request.
*/
func CreateUser(context *gin.Context) {
	var newUser models.User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	Users = append(Users, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

// Unused
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

// Unused
func getUserByID(id uint) (*models.User, error) {
	for i, u := range Users {
		if u.ID == id {
			return &Users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

// Unused
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

	context.IndentedJSON(http.StatusCreated, updatedUser)
}

// Unused
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

// Unused
func removeUser(users []models.User, user *models.User) []models.User {
	for i, u := range users {
		if u.ID == user.ID {
			return append(users[:i], users[i+1:]...)
		}
	}

	return users
}
