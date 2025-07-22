package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

	"trust-credit-back/database"
	"trust-credit-back/models"
	"trust-credit-back/service/security"
	"trust-credit-back/service/utils"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//TODO: рефактор валидации и ручек авторизации\регистрации
//TODO: ручка отправки ID фотографий и удаления фотографий

func InitPhoneValidation(validate *validator.Validate) {
	validate.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
        re := regexp.MustCompile(`^[78][0-9]{10}$`)
        return re.MatchString(fl.Field().String())
    })
}

func InitPasswordValidation(validate *validator.Validate) {
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
        re := regexp.MustCompile(`^(.{0,6}|[^0-9]*|[^A-Z]*|[^a-z]*|[a-zA-Z0-9]*)$`)
        return !re.MatchString(fl.Field().String())
    })
}

func InitBirthDateValidation(validate *validator.Validate) {
	validate.RegisterValidation("date", func(fl validator.FieldLevel) bool {
		_, err := utils.ParseDateOfBirth(fl.Field().String())
		return err == nil
	})
}

func RegUser (c echo.Context) error {
	var (
		auth_cred models.AuthCredentials
		this_phone models.PhoneNumber
		ref models.RegForm
	)

	if err := c.Bind(&ref); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	validate := validator.New()
	InitPasswordValidation(validate)
	InitPhoneValidation(validate)
	InitBirthDateValidation(validate)

	err := validate.Struct(ref)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	found := database.DB.Where("phone_number = ?", ref.PhoneNumber).Find(&this_phone).RowsAffected > 0
	if found {
		return c.JSON(http.StatusConflict, map[string]string{
			"message": "user already exist",
		})
	}

	date, _:= utils.ParseDateOfBirth(ref.DateOfBirth)

	user := models.User{
		ID: uuid.New().String(),
		// AgentUserID: req.AgentUserID,
		FirstName:   ref.FirstName,
		LastName:    ref.LastName,
		DateOfBirth: date,
		AccountType: models.UserType,
		RegDate:     time.Now(),
	}


	if err := database.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to create user",
		})
	}

	if ref.Password == "" {
		auth_cred = models.AuthCredentials{
			AuthType: 	models.PhoneCode,
			Login: 		ref.PhoneNumber,
			UserID: 	user.ID,
		}

		if err := database.DB.Create(&auth_cred).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "failed to create user",
			})
		}
	} else {
		hashed, err := security.GenerateHash(ref.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": "failed to create user",
			})
		}
	
		auth_cred = models.AuthCredentials{
			AuthType: 	models.PhonePassword,
			Login: 		ref.PhoneNumber,
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
		ID:				uuid.New().String(),
		PhoneNumber: 	ref.PhoneNumber,
		UserID:      	user.ID,
	}

	if err := database.DB.Create(&phone).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to create phone number",
		})
	}

	database.DB.Model(&user).Association("PhoneNumbers").Append(&phone)
	database.DB.Model(&user).Association("AuthCredentials").Append(&auth_cred)

	database.DB.Preload("PhoneNumbers").Preload("AuthCredentials").Where("user_id = ?", user.ID).Find(&user)

	tokens, err := security.NewTokens(user.ID)

	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}

func AuthUser (c echo.Context) error {
	login, password := c.FormValue("login"), c.FormValue("password")

	var auth_cred models.AuthCredentials

	found := database.DB.Where("login = ? AND auth_type = ?", login, models.PhonePassword).Find(&auth_cred).RowsAffected > 0
	
	if !found || !security.CompareHash(security.HashedPassword{Salt: auth_cred.Salt, Hash: auth_cred.Hash}, password) {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid credentials",
		})
	}

	tokens, err := security.NewTokens(auth_cred.UserID)

	if err != nil {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err,
		})
	}

	fmt.Println("user.ID:", auth_cred.UserID)
	fmt.Println("user.ID.String():", auth_cred.UserID)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
