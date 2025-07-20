package handlers

import (
	"net/http"
	"trust-credit-back/service"

	"github.com/labstack/echo/v4"
)

func RefreshTokens (c echo.Context) error {
	id, _ := c.Get("id").(string)

 	tokens, err := service.NewTokens(id)

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