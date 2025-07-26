package routes

import (
	"trust-credit-back/handlers"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {
	e.POST("/reg", handlers.RegUser)
	e.POST("/auth", handlers.AuthUser)
	e.POST("/request-code", handlers.RequestCode)
	e.POST("/verify-code", handlers.VerifyCode)
}
