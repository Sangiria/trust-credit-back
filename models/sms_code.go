package models

import "time"

type SMSCode struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Code      string
	CreatedAt time.Time
	ExpiredAt time.Time
	IsUsed    bool
}
