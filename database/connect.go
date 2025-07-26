package database

import (
	"trust-credit-back/environment"
	"trust-credit-back/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db_connection *gorm.DB

func GetDBConnection() *gorm.DB {
	if db_connection == nil {
		connectDB()
	}
	return db_connection
}

func connectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(environment.GetVariable("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	_ = db.AutoMigrate(
		&models.User{},
		&models.PhoneNumber{},
		&models.AuthCredentials{},
		&models.SMSCode{},
	)

	db_connection = db

	return db_connection
}
