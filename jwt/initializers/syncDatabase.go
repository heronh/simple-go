package initializers

import "github.com/heronh/simple-go/jwt/models"

func SyncDatabase() {
	// Sync the database
	DB.AutoMigrate(&models.User{})
}
