package route

import (
	"clean_architecture/config"
	"clean_architecture/controller"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	app := echo.New()
	app.POST("/login", controller.LoginUserController)

	eJwt := app.Group("")
	eJwt.Use(mid.JWT([]byte(config.SECRET_JWT)))
	eJwt.GET("/users", controller.GetAllUsers)
	eJwt.POST("/users", controller.CreateUser)

	return app
}
