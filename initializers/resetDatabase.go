// This is only for simulating all users logged out and no threads in database

package initializers

import "github.com/garylow2001/GossipGo-Backend/models"

func ResetValuesInDatabase() {
	// Delete all tokens to simulate no user logged in
	result := DB.Model(&models.Auth{}).Where("id >= ?", "1").Update("token", nil)
	if result.Error != nil {
		panic(result.Error)
	}
}
