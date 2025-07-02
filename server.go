package main

import (
	"trust-credit-back/database"
	"trust-credit-back/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database.AutoMigrate()

	routes.RegisterUserRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
