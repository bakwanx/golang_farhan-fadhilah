package controller

import (
	"belajar-go-echo/config"
	"belajar-go-echo/model"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {

	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func CreateUser(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	err := config.DB.Create(&user).Error
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": user,
	})
}
