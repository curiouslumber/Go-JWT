package initializers

import (
	"example/go-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
