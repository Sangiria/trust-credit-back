package models

type AuthType string

const (
	PhoneCode AuthType = "phone+code"
	PhonePassword AuthType = "phone+password"
)

type AuthCredentials struct {
	AuthType 	AuthType	`gorm:"primaryKey"`
	Login    	string		`gorm:"primaryKey"`
	Hash	 	string
	Salt     	string
	UserID   	uint
}