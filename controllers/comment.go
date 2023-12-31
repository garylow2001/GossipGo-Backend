package controllers

import (
	"github.com/gin-gonic/gin"
)

// Comment handlers
func CreateComment(context *gin.Context) {
	// TODO: Implement create comment logic
	// threadID := context.Param("threadID")
	// commentID := context.Param("commentID")
}

func GetComment(context *gin.Context) {
	// TODO: Implement get comment logic
	threadID := context.Param("threadID")
	commentID := context.Param("commentID")

	println("threadID: " + threadID)
	println("commentID: " + commentID)

	context.JSON(200, gin.H{
		"threadID":  threadID,
		"commentID": commentID,
	})
}

func UpdateComment(context *gin.Context) {
	// TODO: Implement update comment logic
	// threadID := context.Param("threadID")
	// commentID := context.Param("commentID")
}

func DeleteComment(context *gin.Context) {
	// TODO: Implement delete comment logic
	// threadID := context.Param("threadID")
	// commentID := context.Param("commentID")
}
