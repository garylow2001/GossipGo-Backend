package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/garylow2001/GossipGo-Backend/internal/handlers"
	"github.com/garylow2001/GossipGo-Backend/seed"
)

// var threads []types.Thread
// var comments []types.Comment

func main() {
	router := gin.Default()

	// User endpoints
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)
	router.GET("/users/:id", handlers.GetUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)

	// Thread endpoints
	router.GET("/threads", handlers.GetThreads)
	router.POST("/threads", handlers.CreateThread)
	router.GET("/threads/:id", handlers.GetThread)
	router.PUT("/threads/:id", handlers.UpdateThread)
	router.DELETE("/threads/:id", handlers.DeleteThread)

	// Comment endpoints
	threadGroup := router.Group("/threads/:id")
	threadGroup.POST("/comments", handlers.CreateComment)
	threadGroup.GET("/comments/:id", handlers.GetComment)
	threadGroup.PUT("/comments/:id", handlers.UpdateComment)
	threadGroup.DELETE("/comments/:id", handlers.DeleteComment)

	// Initialize seed data
	handlers.Users = seed.SeededUsers
	// threads = seed.SeededThreads
	// comments = seed.SeededComments

	log.Fatal(http.ListenAndServe(":8080", router))
}
