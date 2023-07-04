package route

import (
	"orm_and_mvc/constant"
	"orm_and_mvc/controller"
	m "orm_and_mvc/middleware"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	m.LogMiddleware(e)
	e.POST("/login", controller.LoginUserController)
	e.POST("/users", controller.CreateUserController)
	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)

	eJwt := e.Group("")
	eJwt.Use(mid.JWT([]byte(constant.SECRET_JWT)))
	eJwt.GET("/users", controller.GetUsersController)
	eJwt.GET("/users/:id", controller.GetUserController)
	eJwt.PUT("/users/:id", controller.UpdateUserController)
	eJwt.DELETE("/users/:id", controller.DeleteUserController)
	eJwt.POST("/books", controller.CreateBookController)
	eJwt.PUT("/books/:id", controller.UpdateBookController)
	eJwt.DELETE("/books/:id", controller.DeleteBookController)
	return e
}
