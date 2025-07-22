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
	UserID   	string
}

type RegForm struct {
	// AgentUserID uint   `json:"agent_user_id" validate:"required"` - убрала на время, пока поле не используется
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required,date"`
	PhoneNumber string `json:"phone_number" validate:"phone"`
	Password	string `json:"password" validate:"omitempty,password"` //убрать????
}