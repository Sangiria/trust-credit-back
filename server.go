package main

import (
	"log"
	"net/http"
	custommiddleware "trust-credit-back/custom_middleware"
	"trust-credit-back/environment"
	"trust-credit-back/routes"
	"trust-credit-back/service"

	"github.com/labstack/echo/v4"
)

func main() {

	if err := service.InitTelegramBot(environment.GetVariable("TELEGRAM_TOKEN")); err != nil {
		log.Fatal("Telegram init failed:", err)
	}

	e := echo.New()

	access := e.Group("/access", custommiddleware.JWTMiddleware(environment.GetVariable("ACCESS_SECRET")))
	refresh := e.Group("/refresh", custommiddleware.JWTMiddleware(environment.GetVariable("REFRESH_SECRET")))

	access.GET("/test", func(c echo.Context) error {
		id, _ := c.Get("id").(string)
		return c.JSON(http.StatusOK, map[string]string{
			"ID": id,
		})
	})

	routes.InitUserRoutes(e)
	routes.InitRefreshRoute(refresh)

	e.Logger.Fatal(e.Start(":1323"))
}
