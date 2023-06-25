package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func GetUsersController(c echo.Context) error {
	if users == nil {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "users empty",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetUserController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	var user User
	var isFound bool
	for _, v := range users {
		if v.Id == id {
			user = v
			isFound = true
		}
	}

	if isFound {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "success get user detail",
			"user":     user,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

func DeleteUserController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	var tempUser []User

	for _, v := range users {
		if v.Id != id {
			tempUser = append(tempUser, v)
		}
	}

	if len(tempUser) == len(users) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "user not found",
		})
	}

	users = tempUser

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success delete user",
		"users":    users,
	})
}

func UpdateUserController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))
	var isFound bool
	var tempUser []User
	var user = User{}
	c.Bind(&user)

	for _, v := range users {
		if v.Id == id {
			v = user
			v.Id = id
			tempUser = append(tempUser, v)
			isFound = true
		} else {
			tempUser = append(tempUser, v)
		}
	}

	users = tempUser

	if isFound {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "success update user",
			"user":     tempUser,
		})
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.DELETE("/users/:id", DeleteUserController)
	e.PUT("/users/:id", UpdateUserController)

	e.Logger.Fatal(e.Start(":8000"))
}
