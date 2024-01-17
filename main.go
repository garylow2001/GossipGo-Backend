package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"

	"github.com/garylow2001/GossipGo-Backend/configs"
	"github.com/garylow2001/GossipGo-Backend/controllers"
	"github.com/garylow2001/GossipGo-Backend/initializers"
	"github.com/garylow2001/GossipGo-Backend/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	setUpCORS(router)
	setUpRouters(router)

	router.Run()
}

func setUpCORS(router *gin.Engine) {
	c := cors.New(cors.Options{
		AllowedOrigins:   configs.CORSAllowedOrigins,
		AllowedMethods:   configs.CORSAllowedMethods,
		AllowedHeaders:   configs.CORSAllowedHeaders,
		AllowCredentials: true, //this is to allow cookies to be included in requests from frontend
	})

	router.Use(func(context *gin.Context) {
		c.HandlerFunc(context.Writer, context.Request)

		// CORS preflight check
		if context.Request.Method != "OPTIONS" {
			context.Next()
		} else {
			context.AbortWithStatus(http.StatusOK)
		}
	})
}

func setUpRouters(router *gin.Engine) {
	// User endpoints
	router.GET("/validate", middleware.JWTAuthMiddleware, controllers.Validate)
	router.GET("/users", middleware.JWTAuthMiddleware, controllers.GetUsers)
	router.POST("/users/signup", controllers.Signup)
	router.POST("/users/login", controllers.Login)
	router.POST("/users/logout", controllers.Logout)
	router.GET("/users/:id", controllers.GetUser)       // Not in use
	router.PUT("/users/:id", controllers.UpdateUser)    // Not in use
	router.DELETE("/users/:id", controllers.DeleteUser) // Not in use

	// Thread endpoints
	router.GET("/threads", controllers.GetThreads)
	router.GET("/threads/category/:category", controllers.GetThreadsByCategory)
	router.GET("/threads/recent", controllers.GetThreadsByMostRecent)
	router.GET("/threads/popular", controllers.GetThreadsByMostPopular)
	router.POST("/threads", middleware.JWTAuthMiddleware, controllers.CreateThread)
	router.POST("/threads/:threadID/like", middleware.JWTAuthMiddleware, controllers.LikeThread)
	router.DELETE("/threads/:threadID/like", middleware.JWTAuthMiddleware, controllers.UnlikeThread)
	router.GET("/threads/:threadID", controllers.GetThread)
	router.PUT("/threads/:threadID", middleware.JWTAuthMiddleware, controllers.UpdateThread)
	router.DELETE("/threads/:threadID", middleware.JWTAuthMiddleware, controllers.DeleteThread)

	// Comment endpoints
	threadGroup := router.Group("/threads/:threadID")
	threadGroup.POST("/comments", middleware.JWTAuthMiddleware, controllers.CreateComment)
	threadGroup.POST("/comments/:commentID/like", middleware.JWTAuthMiddleware, controllers.LikeComment)
	threadGroup.DELETE("/comments/:commentID/like", middleware.JWTAuthMiddleware, controllers.UnlikeComment)
	threadGroup.GET("/comments", controllers.GetComments)
	threadGroup.GET("/comments/:commentID", controllers.GetComment)
	threadGroup.PUT("/comments/:commentID", middleware.JWTAuthMiddleware, controllers.UpdateComment)
	threadGroup.DELETE("/comments/:commentID", middleware.JWTAuthMiddleware, controllers.DeleteComment)
}
