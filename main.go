package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/garylow2001/GossipGo-Backend/seed"
)

func main() {
	router := gin.Default()

	// User endpoints
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
	users := seed.SeededUsers
	threads := seed.SeededThreads
	comments := seed.SeededComments

	log.Fatal(http.ListenAndServe(":8080", router))
}

// User handlers
func createUser(c *gin.Context) {
	// TODO: Implement create user logic
}

func getUser(c *gin.Context) {
	// TODO: Implement get user logic
}

func updateUser(c *gin.Context) {
	// TODO: Implement update user logic
}

func deleteUser(c *gin.Context) {
	// TODO: Implement delete user logic
}

// Thread handlers
func createThread(c *gin.Context) {
	// TODO: Implement create thread logic
}

func getThread(c *gin.Context) {
	// TODO: Implement get thread logic
}

func updateThread(c *gin.Context) {
	// TODO: Implement update thread logic
}

func deleteThread(c *gin.Context) {
	// TODO: Implement delete thread logic
}

// Comment handlers
func createComment(c *gin.Context) {
	// TODO: Implement create comment logic
	threadID := c.Param("id")
}

func getComment(c *gin.Context) {
	// TODO: Implement get comment logic
	threadID := c.Param("id")
	commentID := c.Param("commentID")
}

func updateComment(c *gin.Context) {
	// TODO: Implement update comment logic
	threadID := c.Param("id")
	commentID := c.Param("commentID")
}

func deleteComment(c *gin.Context) {
	// TODO: Implement delete comment logic
	threadID := c.Param("id")
	commentID := c.Param("commentID")
}
