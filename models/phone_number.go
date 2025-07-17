package models

type PhoneNumber struct {
	ID          string `gorm:"primaryKey"`
	PhoneNumber string
	UserID      string
}
