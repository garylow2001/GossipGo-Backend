package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/garylow2001/GossipGo-Backend/models"
	"github.com/gin-gonic/gin"
)

var Threads []models.Thread

// Thread handlers
func GetThreads(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Threads)
}

func CreateThread(context *gin.Context) {
	var newThread models.Thread

	// TODO: validate author
	if err := context.BindJSON(&newThread); err != nil {
		return
	}

	Threads = append(Threads, newThread)

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

func getThreadByID(id int) (*models.Thread, error) {
	for i, t := range Threads {
		if t.ID == id {
			return &Threads[i], nil
		}
	}

	return nil, errors.New("thread not found")
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

	Threads = removeThread(Threads, thread)

	context.IndentedJSON(http.StatusOK, thread)
}

func removeThread(threads []models.Thread, thread *models.Thread) []models.Thread {
	for i, t := range threads {
		if t.ID == thread.ID {
			threads = append(threads[:i], threads[i+1:]...)
			break
		}
	}

	return threads
}
