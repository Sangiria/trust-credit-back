package database

import (
	"trust-credit-back/models"
)

func AutoMigrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.PhoneNumber{},
	)
}
