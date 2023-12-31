package initializers

import (
	"github.com/garylow2001/GossipGo-Backend/models"
)

func SyncDatabase() {
	// Auto migrate all models
	err := DB.AutoMigrate(&models.User{}, &models.Auth{}, &models.Thread{}, &models.Comment{})
	if err != nil {
		panic(err)
	}
}
