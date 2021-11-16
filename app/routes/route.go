package route

import (
	controller "widi443/currency-exchange/app/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controller.GetUserController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/wallet", controller.GetWallet)
	e.POST("/wallet", controller.CreateWallet)

	return e

}