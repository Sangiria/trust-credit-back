package main

import (
	"trust-credit-back/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.InitUserRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
