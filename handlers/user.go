package handlers

import (
	"net/http"
	"time"

	"trust-credit-back/database"
	"trust-credit-back/models"
	"trust-credit-back/service"

	"github.com/labstack/echo/v4"
)

type RegUserRequest struct {
	AgentUserID uint   `json:"agent_user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	AccountType string `json:"account_type"`
	PhoneNumber string `json:"phone_number"`
	Password	string `json:"password"`
}

func RegUser (c echo.Context) error {
	var (
		auth_cred models.AuthCredentials
		this_phone models.PhoneNumber
		req RegUserRequest
	)

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	database.DB.Where("phone_number = ?", req.PhoneNumber).Find(&this_phone)
	if this_phone.ID != 0 {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": "user already exist",
		})
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
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to create user",
		})
	}


	if req.Password == "" {
		auth_cred = models.AuthCredentials{
			AuthType: 	models.PhoneCode,
			Login: 		req.PhoneNumber,
			UserID: 	user.ID,
		}

		if err := database.DB.Create(&auth_cred).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "failed to create user",
			})
		}
	} else {
		hashed, err := service.GenerateHash(req.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "failed to create user",
			})
		}
	
		auth_cred = models.AuthCredentials{
			AuthType: 	models.PhonePassword,
			Login: 		req.PhoneNumber,
			Salt:		hashed.Salt,
			Hash:		hashed.Hash,
			UserID: 	user.ID,
		}
	
		if err := database.DB.Create(&auth_cred).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "failed to create user",
			})
		}
	}

	phone := models.PhoneNumber{
		PhoneNumber: req.PhoneNumber,
		UserID:      user.ID,
	}

	if err := database.DB.Create(&phone).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to create phone number",
		})
	}

	database.DB.Model(&user).Association("PhoneNumbers").Append(&phone)
	database.DB.Model(&user).Association("AuthCredentials").Append(&auth_cred)

	database.DB.Preload("PhoneNumbers").Preload("AuthCredentials").Where("user_id = ?", user.ID).Find(&user)

	return c.JSON(http.StatusOK, user)
}

func AuthUser (c echo.Context) error {
	login, password := c.FormValue("login"), c.FormValue("password")

	var auth_cred models.AuthCredentials

	found := database.DB.Where("login = ? AND auth_type = ?", login, models.PhonePassword).Find(&auth_cred).RowsAffected > 0

	if !found || !service.CompareHash(service.HashedPassword{Salt: auth_cred.Salt, Hash: auth_cred.Hash}, password) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid credentials",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})

}
