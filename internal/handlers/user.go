package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/garylow2001/GossipGo-Backend/types"
	"github.com/gin-gonic/gin"
)

var Users []types.User

// User handlers
func GetUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Users)
}

func CreateUser(context *gin.Context) {
	var newUser types.User

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

	user, err := getUserByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"}) //TODO: abstract out error message
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func getUserByID(id int) (*types.User, error) {
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

	user, err := getUserByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	var updatedUser types.User

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

	user, err := getUserByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	Users = removeUser(Users, user)

	context.IndentedJSON(http.StatusCreated, Users)
}

func removeUser(Users []types.User, user *types.User) []types.User {
	for i, u := range Users {
		if u.ID == user.ID {
			return append(Users[:i], Users[i+1:]...)
		}
	}

	return Users
}
