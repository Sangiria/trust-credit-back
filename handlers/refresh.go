package handlers

import (
	"net/http"
	"trust-credit-back/service"

	"github.com/labstack/echo/v4"
)

func RefreshTokens (c echo.Context) error {
	id, _ := c.Get("id").(string)

 	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "0",
		"tokens": map[string]string{
			"access_token": service.NewToken(id, true),
			"refresh_token": service.NewToken(id, false),
		},
	})
}