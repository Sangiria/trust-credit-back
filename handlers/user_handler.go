package handlers

import (
	"net/http"
	"time"

	"trust-credit-back/database"
	"trust-credit-back/models"

	"github.com/labstack/echo/v4"
)

type CreateUserRequest struct {
	AgentUserID uint   `json:"agent_user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	AccountType string `json:"account_type"`
	PhoneNumber string `json:"phone_number"`
}

func CreateUser(c echo.Context) error {
	var req CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// Создаём пользователя
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

	// Создаём телефон
	phone := models.PhoneNumber{
		PhoneNumber: req.PhoneNumber,
		UserID:      user.ID,
	}
	if err := database.DB.Create(&phone).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to create phone number"})
	}

	// Подгружаем связанные телефоны
	database.DB.Preload("PhoneNumbers").First(&user, user.ID)

	// Возвращаем созданного пользователя
	return c.JSON(http.StatusOK, user)
}
