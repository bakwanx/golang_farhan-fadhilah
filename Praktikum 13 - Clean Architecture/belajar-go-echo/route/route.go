package route

import (
	"belajar-go-echo/config"

	"belajar-go-echo/controller"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(config.SECRET_JWT)))
	eJwt.GET("/users", controller.GetAllUsers(db))
	e.POST("/users", controller.CreateUser(db))

	return e
}
