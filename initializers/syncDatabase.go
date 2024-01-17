package initializers

import (
	"github.com/garylow2001/GossipGo-Backend/models"
)

func SyncDatabase() {
	// Auto migrate all models
	err := DB.AutoMigrate(&models.User{}, &models.Auth{}, &models.Thread{}, &models.Comment{}, models.CommentLike{}, models.ThreadLike{})
	if err != nil {
		panic(err)
	}

	// Ensure the likes_count tally with the actual number of likes in the database
	var threads []models.Thread
	result := DB.Preload("Likes").Find(&threads)
	if result.Error != nil {
		panic(result.Error)
	}

	for _, thread := range threads {
		likesCount := len(thread.Likes)
		thread.LikesCount = likesCount
		err := DB.Model(thread).Update("LikesCount", thread.LikesCount).Error
		if err != nil {
			panic(err)
		}
	}
}
