package models

type AuthCredentials struct {
	AuthType string	`gorm:"primaryKey"`
	Login    string	`gorm:"primaryKey"`
	Hash	 string
	Salt     string
	//ExpDate ?
	UserID   uint
}