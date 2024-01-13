package controllers

import (
	"net/http"
	"strconv"

	"github.com/garylow2001/GossipGo-Backend/initializers"
	"github.com/garylow2001/GossipGo-Backend/models"
	"github.com/garylow2001/GossipGo-Backend/utils"
	"github.com/gin-gonic/gin"
)

// Thread handlers
func GetThreads(context *gin.Context) {
	var threads []models.Thread

	result := initializers.DB.Preload("Author").Find(&threads) // Comments not preloaded as it is not needed here
	if result.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Error retrieving threads"})
		return
	}

	context.IndentedJSON(http.StatusOK, threads)
}

func GetThreadsByCategory(context *gin.Context) {
	var threads []models.Thread

	category := context.Param("category")

	if !isValidCategory(category) {
		category = utils.CapitalizeFirstLetter(category)
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	result := initializers.DB.Preload("Author").Where("category = ?", category).Find(&threads) // Comments not preloaded as it is not needed here
	if result.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Error retrieving threads"})
		return
	}

	context.IndentedJSON(http.StatusOK, threads)
}

func GetThreadsByMostRecent(context *gin.Context) {
	var threads []models.Thread

	result := initializers.DB.Preload("Author").Order("updated_at desc").Find(&threads) // Comments not preloaded as it is not needed here
	if result.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Error retrieving threads"})
		return
	}

	context.IndentedJSON(http.StatusOK, threads)
}

func CreateThread(context *gin.Context) {
	var newThread models.Thread

	user := context.MustGet("user").(models.User)

	if err := context.ShouldBindJSON(&newThread); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Check if category is valid
	checkValidCategory(&newThread, context)

	// Tag the thread with the author
	newThread.Author = user

	// Add new thread to the database
	result := initializers.DB.Create(&newThread)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	context.IndentedJSON(http.StatusCreated, newThread)
}

func GetThread(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("threadID"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"}) //TODO: abstract out invalid integer error message
		return
	}

	thread, err := getThreadByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Thread not found"}) //TODO: abstract out error message
		return
	}

	context.IndentedJSON(http.StatusOK, thread)
}

func UpdateThread(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("threadID"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"}) //TODO: abstract out invalid integer error message
		return
	}

	thread, err := getThreadByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Thread not found"}) //TODO: abstract out error message
		return
	}

	// check if user is author of thread
	user := context.MustGet("user").(models.User)

	if thread.AuthorID != user.ID {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. You are not the author of this thread"})
		return
	}

	// Update thread
	var updatedThread models.Thread

	if err := context.BindJSON(&updatedThread); err != nil {
		return
	}

	// Check if category is valid
	checkValidCategory(&updatedThread, context)

	thread.Title = updatedThread.Title
	thread.Body = updatedThread.Body

	initializers.DB.Save(&thread)

	context.IndentedJSON(http.StatusOK, thread)
}

func DeleteThread(context *gin.Context) {
	// TODO: ensure only author can delete thread
	id, err := strconv.Atoi(context.Param("threadID"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"}) //TODO: abstract out invalid integer error message
		return
	}

	thread, err := getThreadByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Thread not found"}) //TODO: abstract out error message
		return
	}

	// check if user is author of thread
	user := context.MustGet("user").(models.User)

	if thread.AuthorID != user.ID {
		context.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. You are not the author of this thread"})
		return
	}

	// Delete comments under thread
	var comments []models.Comment
	initializers.DB.Where("thread_id = ?", thread.ID).Find(&comments)

	if len(comments) > 0 {
		result := initializers.DB.Delete(&comments)
		if result.Error != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comments under the thread"})
			return
		}
	}

	// Delete thread
	initializers.DB.Delete(&thread, id)

	context.IndentedJSON(http.StatusOK, thread)
}

func getThreadByID(id int) (*models.Thread, error) {
	var thread models.Thread

	result := initializers.DB.Preload("Author").First(&thread, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &thread, nil
}

func checkValidCategory(thread *models.Thread, context *gin.Context) {
	if thread.Category != "" && !isValidCategory(thread.Category) {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		context.Abort()
	}
}

func isValidCategory(category string) bool {
	for _, validCategory := range models.ValidCategories {
		if category == validCategory {
			return true
		}
	}

	return false
}
