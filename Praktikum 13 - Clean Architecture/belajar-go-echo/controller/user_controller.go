package controller

import (
	"belajar-go-echo/model"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DB interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
}

func GetAllUsers(db DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		users := make([]model.User, 0)
		err := db.Find(&users).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": users,
		})
	}
}

func CreateUser(db DB) echo.HandlerFunc {
	user := model.User{}
	return func(c echo.Context) error {
		if err := c.Bind(&user); err != nil {
			return c.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}
		err := db.Create(&user).Error
		if err != nil {
			return c.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(200, echo.Map{
			"data": user,
		})
	}
}
