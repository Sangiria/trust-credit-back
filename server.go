package main

import (
	"net/http"
	custommiddleware "trust-credit-back/custom_middleware"
	"trust-credit-back/environment"
	"trust-credit-back/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	access := e.Group("/access", custommiddleware.JWTMiddleware(environment.GetVariable("ACCESS_SECRET")))
	refresh := e.Group("/refresh", custommiddleware.JWTMiddleware(environment.GetVariable("REFRESH_SECRET")))

	access.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "test has been successful.")
	})

	routes.InitUserRoutes(e)
	routes.InitRefreshRoute(refresh)
	
	e.Logger.Fatal(e.Start(":1323"))
}
