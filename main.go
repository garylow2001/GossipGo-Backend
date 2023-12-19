package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/garylow2001/GossipGo-Backend/seed"
	"github.com/garylow2001/GossipGo-Backend/types"
)

var users []types.User

// var threads []types.Thread
// var comments []types.Comment

func main() {
	router := gin.Default()

	// User endpoints
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUser)
	router.PUT("/users/:id", updateUser)
	router.DELETE("/users/:id", deleteUser)

	// Thread endpoints
	router.POST("/threads", createThread)
	router.GET("/threads/:id", getThread)
	router.PUT("/threads/:id", updateThread)
	router.DELETE("/threads/:id", deleteThread)

	// Comment endpoints
	threadGroup := router.Group("/threads/:id")
	threadGroup.POST("/comments", createComment)
	threadGroup.GET("/comments/:id", getComment)
	threadGroup.PUT("/comments/:id", updateComment)
	threadGroup.DELETE("/comments/:id", deleteComment)

	// Initialize seed data
	users = seed.SeededUsers
	// threads = seed.SeededThreads
	// comments = seed.SeededComments

	log.Fatal(http.ListenAndServe(":8080", router))
}

// User handlers
func getUsers(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, users)
}

func createUser(context *gin.Context) {
	var newUser types.User

	if err := context.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)

	context.IndentedJSON(http.StatusCreated, newUser)
}

func getUser(context *gin.Context) {
	id := context.Param("id")
	user, err := getUserByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"}) //TODO: abstract out error message
		return
	}

	context.IndentedJSON(http.StatusOK, user)
}

func getUserByID(id string) (*types.User, error) {
	for i, t := range users {
		if t.ID == id {
			return &users[i], nil
		}
	}

	return nil, errors.New("user not found")
}

func updateUser(context *gin.Context) {
	// TODO: Implement update user logic
	id := context.Param("id")
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

func deleteUser(context *gin.Context) {
	// TODO: Implement delete user logic
}

// Thread handlers
func createThread(context *gin.Context) {
	// TODO: Implement create thread logic
}

func getThread(context *gin.Context) {
	// TODO: Implement get thread logic
}

func updateThread(context *gin.Context) {
	// TODO: Implement update thread logic
}

func deleteThread(context *gin.Context) {
	// TODO: Implement delete thread logic
}

// Comment handlers
func createComment(context *gin.Context) {
	// TODO: Implement create comment logic
	// threadID := c.Param("id")
}

func getComment(context *gin.Context) {
	// TODO: Implement get comment logic
	// threadID := c.Param("id")
	// commentID := c.Param("commentID")
}

func updateComment(context *gin.Context) {
	// TODO: Implement update comment logic
	// threadID := c.Param("id")
	// commentID := c.Param("commentID")
}

func deleteComment(context *gin.Context) {
	// TODO: Implement delete comment logic
	// threadID := c.Param("id")
	// commentID := c.Param("commentID")
}
