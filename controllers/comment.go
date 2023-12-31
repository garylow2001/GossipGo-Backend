package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/garylow2001/GossipGo-Backend/initializers"
	"github.com/garylow2001/GossipGo-Backend/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Comment handlers

func GetComments(context *gin.Context) {
	// TODO: Sort by most recently updated/created
}

func CreateComment(context *gin.Context) {
	var newComment models.Comment

	if err := context.ShouldBindJSON(&newComment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	user := context.MustGet("user").(models.User)
	threadIDString := context.Param("threadID")
	threadID, err := strconv.ParseUint(threadIDString, 10, strconv.IntSize)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threadID"})
		return
	}

	commentID, err := getLastCommentID(uint(threadID))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve last commentID"})
		return
	}

	// Tag the comment with the author and thread, and set the commentID
	newComment.CommentID = commentID + 1
	newComment.Author = user
	newComment.ThreadID = uint(threadID) // might need to change to newComment.Thread = thread instead

	// Add new comment to the database
	result := initializers.DB.Create(&newComment)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	context.IndentedJSON(http.StatusCreated, newComment)
}

func GetComment(context *gin.Context) {
	comment, err := retrieveComment(context)
	if err != nil {
		return
	}

	context.IndentedJSON(http.StatusOK, comment)
}

func UpdateComment(context *gin.Context) {
	comment, err := retrieveComment(context)
	if err != nil {
		return
	}

	var updatedComment models.Comment
	if err := context.ShouldBindJSON(&updatedComment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Update the comment
	result := initializers.DB.Model(&comment).Updates(updatedComment)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	context.IndentedJSON(http.StatusOK, comment)
}

func DeleteComment(context *gin.Context) {
	comment, err := retrieveComment(context)
	if err != nil {
		return
	}

	result := initializers.DB.Delete(&comment)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	context.IndentedJSON(http.StatusOK, comment)
}

func retrieveComment(context *gin.Context) (*models.Comment, error) {
	threadIDString := context.Param("threadID")
	threadID, err := strconv.ParseUint(threadIDString, 10, strconv.IntSize)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid threadID"})
		return nil, err
	}

	commentIDString := context.Param("commentID")
	commentID, err := strconv.ParseUint(commentIDString, 10, strconv.IntSize)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid commentID"})
		return nil, err
	}

	var comment models.Comment
	result := initializers.DB.Where("thread_id = ? AND comment_id = ?", threadID, commentID).First(&comment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			context.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
			return nil, err
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comment"})
			return nil, err
		}
	}

	return &comment, nil
}

func getLastCommentID(threadID uint) (uint, error) {
	var lastComment models.Comment
	result := initializers.DB.Where("thread_id = ?", threadID).Order("comment_id").Last(&lastComment)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil
		} else {
			return 0, result.Error
		}
	}

	return lastComment.CommentID, nil
}
