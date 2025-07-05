package models

type AuthCredentials struct {
	AuthType 	AuthType	`gorm:"primaryKey"`
	Login    	string		`gorm:"primaryKey"`
	Hash	 	string
	Salt     	string
	//ExpDate ?
	UserID   	uint
}