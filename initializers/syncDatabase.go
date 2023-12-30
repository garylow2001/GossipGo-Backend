package initializers

import (
	"github.com/garylow2001/GossipGo-Backend/models"
)

func SyncDatabase() {
	// Reset threads table
	err := DB.Migrator().DropTable(&models.Thread{})
	if err != nil {
		panic(err)
	}

	// Auto migrate all models
	err = DB.AutoMigrate(&models.User{}, &models.Auth{}, &models.Thread{})
	if err != nil {
		panic(err)
	}
}
