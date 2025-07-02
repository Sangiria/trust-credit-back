package models

type PhoneNumber struct {
	ID          uint `gorm:"primaryKey"`
	PhoneNumber string
	UserID      uint
}
