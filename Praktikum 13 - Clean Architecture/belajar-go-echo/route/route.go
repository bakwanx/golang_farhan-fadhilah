package route

import (
	"belajar-go-echo/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	app := echo.New()
	app.GET("/users", controller.GetAllUsers)
	app.POST("/users", controller.CreateUser)

	return app
}
