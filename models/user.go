package models

import "time"

type User struct {
	ID           	uint 				`gorm:"primaryKey"`
	// AgentUserID  	uint
	FirstName    	string
	LastName     	string
	MiddleName   	string
	RegDate      	time.Time
	AccountType  	string
	PhoneNumbers 	[]PhoneNumber 		`gorm:"foreignKey:UserID"`
	AuthCredentials []AuthCredentials	`gorm:"foreignKey:UserID"`
}
