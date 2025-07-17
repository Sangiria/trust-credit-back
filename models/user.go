package models

import "time"

type AccountType string

const (
	UserType = "user"
)

type User struct {
	ID           	string 				`gorm:"primaryKey"`
	//AgentUserID  	uint
	FirstName    	string
	LastName     	string
	DateOfBirth     time.Time 			//как будет передаваться значение, как его валидировать и нужно ли вообще?
	RegDate      	time.Time			//как хранить значения, только дату (дд.мм.гг) или еще + время?
	AccountType  	AccountType
	PhoneNumbers 	[]PhoneNumber 		`gorm:"foreignKey:UserID"`
	AuthCredentials []AuthCredentials	`gorm:"foreignKey:UserID"`
}
