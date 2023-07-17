package controller

import (
	"net/http"
	"restfull_test/config"
	"restfull_test/middleware"
	"restfull_test/model"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	var users []model.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success all users",
		"data":    users,
	})
}

func GetUserController(c echo.Context) error {
	id := c.Param("id")
	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user":    user,
	})
}

func CreateUserController(c echo.Context) error {

	user := model.User{}
	c.Bind(&user)

	if user.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"message": "password required",
		})
	}

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	var user model.User

	if err := config.DB.Delete(&user, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete",
	})
}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	var user model.User

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&user)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by id",
		"user":    user,
	})
}

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	if err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	token, err := middleware.CreateToken(user.Id, user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "fail login",
			"error":   err.Error(),
		})
	}

	userResponse := model.UserResponse{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    userResponse,
	})
}
