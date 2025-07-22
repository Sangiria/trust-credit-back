package routes

import (
	"trust-credit-back/handlers"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {
	e.POST("/reg/qr-code", handlers.SendFormInQR)
	e.POST("/reg", handlers.RegUser)
	e.POST("/auth", handlers.AuthUser)
}
