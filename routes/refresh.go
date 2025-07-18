package routes

import (
	"trust-credit-back/handlers"

	"github.com/labstack/echo/v4"
)

func InitRefreshRoute(refresh *echo.Group) {
	refresh.POST("", handlers.RefreshTokens)
}