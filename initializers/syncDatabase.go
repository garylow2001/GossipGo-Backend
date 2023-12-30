package initializers

import (
	"github.com/garylow2001/GossipGo-Backend/models"
)

func SyncDatabase() {
	err := DB.AutoMigrate(&models.User{}, &models.Auth{})
	if err != nil {
		panic(err)
	}
}
