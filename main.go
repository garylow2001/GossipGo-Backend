package main

import (
	"github.com/gin-gonic/gin"

	"github.com/garylow2001/GossipGo-Backend/controllers"
	"github.com/garylow2001/GossipGo-Backend/initializers"
	"github.com/garylow2001/GossipGo-Backend/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
	initializers.ResetValuesInDatabase()
}

func main() {
	router := gin.Default()

	setUpRouters(router)

	router.Run()
}

func setUpRouters(router *gin.Engine) {
	// User endpoints
	router.GET("/validate", middleware.JWTAuthMiddleware, controllers.Validate) //testing with authmiddleware
	router.GET("/users", middleware.JWTAuthMiddleware, controllers.GetUsers)
	router.POST("/users/signup", controllers.Signup)
	router.POST("/users/login", controllers.Login)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// Thread endpoints
	router.GET("/threads", controllers.GetThreads)
	router.POST("/threads", controllers.CreateThread)
	router.GET("/threads/:id", controllers.GetThread)
	router.PUT("/threads/:id", controllers.UpdateThread)
	router.DELETE("/threads/:id", controllers.DeleteThread)

	// Comment endpoints
	threadGroup := router.Group("/threads/:id")
	threadGroup.POST("/comments", controllers.CreateComment)
	threadGroup.GET("/comments/:id", controllers.GetComment)
	threadGroup.PUT("/comments/:id", controllers.UpdateComment)
	threadGroup.DELETE("/comments/:id", controllers.DeleteComment)
}
