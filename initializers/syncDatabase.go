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
	var comments []models.Comment
	result := DB.Preload("Likes").Find(&comments)
	if result.Error != nil {
		panic(result.Error)
	}

	for i := range comments {
		likesCount := len(comments[i].Likes)
		comments[i].LikesCount = likesCount
		err := DB.Model(&comments[i]).Update("LikesCount", comments[i].LikesCount).Error
		if err != nil {
			panic(err)
		}
	}
}
