package routes

import (
	"trust-credit-back/handlers"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo) {
	e.POST("/users", handlers.CreateUser)
}
