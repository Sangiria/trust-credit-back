package models

type AuthType string

const (
	PhoneCode AuthType = "phone+code"
	PhonePassword AuthType = "phone+password"
)