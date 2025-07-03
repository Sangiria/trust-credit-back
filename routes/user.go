package routes

import (
	"trust-credit-back/handlers"

	"github.com/labstack/echo/v4"
)

func InitUserRoutes(e *echo.Echo) {
	e.POST("/users", handlers.CreateUser)
}
