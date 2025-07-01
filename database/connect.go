package database

import (
	"trust-credit-back/environment"

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

func connectDB () *gorm.DB {
	db, err := gorm.Open(postgres.Open(environment.GetVariable("DATABASE_URL")), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database")
	}

	db_connection = db

	return db_connection
}