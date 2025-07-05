package handlers

import (
	"net/http"
	"strings"
	"time"

	"trust-credit-back/database"
	"trust-credit-back/models"
	"trust-credit-back/service"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	AgentUserID uint   `json:"agent_user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	AccountType string `json:"account_type"`
	PhoneNumber string `json:"phone_number"`
	Password	string `json:"password"`
}

func CreateUser(c echo.Context) error {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	user := models.User{
		AgentUserID: req.AgentUserID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		MiddleName:  req.MiddleName,
		AccountType: req.AccountType,
		RegDate:     time.Now(),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create user"})
	}

	var auth_cred models.AuthCredentials

	if req.Password == "" {
		auth_cred = models.AuthCredentials{
			AuthType: 	models.PhoneCode,
			Login: 		req.PhoneNumber,
			UserID: 	user.ID,
		}

		if err := database.DB.Create(&auth_cred).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create user"})
		}
	} else {
		salt_hash := strings.Split(service.GenerateHash(req.Password), "&")
	
		auth_cred = models.AuthCredentials{
			AuthType: 	models.PhonePassword,
			Login: 		req.PhoneNumber,
			Salt:		salt_hash[0],
			Hash:		salt_hash[1],
			UserID: 	user.ID,
		}
	
		if err := database.DB.Create(&auth_cred).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create user"})
		}
	}

	phone := models.PhoneNumber{
		PhoneNumber: req.PhoneNumber,
		UserID:      user.ID,
	}

	if err := database.DB.Create(&phone).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create phone number"})
	}

	database.DB.Model(&user).Association("PhoneNumbers").Append(&phone)
	database.DB.Model(&user).Association("AuthCredentials").Append(&auth_cred)

	database.DB.Preload("PhoneNumbers").Preload("AuthCredentials").Where("user_id = ?", user.ID).Find(&user)

	return c.JSON(http.StatusOK, user)
}
