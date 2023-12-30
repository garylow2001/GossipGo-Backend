package controllers

import (
	"net/http"
	"strconv"

	"github.com/garylow2001/GossipGo-Backend/initializers"
	"github.com/garylow2001/GossipGo-Backend/models"
	"github.com/gin-gonic/gin"
)

// Thread handlers
func GetThreads(context *gin.Context) {
	var threads []models.Thread
	result := initializers.DB.Find(&threads)

	if result.Error != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Error retrieving threads"})
		return
	}

	if len(threads) == 0 {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "No threads found"})
		return
	}

	context.IndentedJSON(http.StatusOK, threads)
}

func CreateThread(context *gin.Context) {
	var newThread models.Thread

	// TODO: validate author
	if err := context.BindJSON(&newThread); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Add new thread to the database
	result := initializers.DB.Create(&newThread)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create thread"})
		return
	}

	context.IndentedJSON(http.StatusCreated, newThread)
}

func GetThread(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"}) //TODO: abstract out invalid integer error message
		return
	}

	thread, err := getThreadByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Thread not found"}) //TODO: abstract out error message
		return
	}

	context.IndentedJSON(http.StatusOK, thread)
}

func UpdateThread(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"}) //TODO: abstract out invalid integer error message
		return
	}

	thread, err := getThreadByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Thread not found"}) //TODO: abstract out error message
		return
	}

	var updatedThread models.Thread

	if err := context.BindJSON(&updatedThread); err != nil {
		return
	}

	thread.Title = updatedThread.Title
	thread.Body = updatedThread.Body

	initializers.DB.Save(&thread)

	context.IndentedJSON(http.StatusOK, thread)
}

func DeleteThread(context *gin.Context) {
	// TODO: ensure only author can delete thread
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"}) //TODO: abstract out invalid integer error message
		return
	}

	thread, err := getThreadByID(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Thread not found"}) //TODO: abstract out error message
		return
	}

	initializers.DB.Delete(&thread, id)

	context.IndentedJSON(http.StatusOK, thread)
}

func getThreadByID(id int) (*models.Thread, error) {
	var thread models.Thread

	result := initializers.DB.First(&thread, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &thread, nil
}
