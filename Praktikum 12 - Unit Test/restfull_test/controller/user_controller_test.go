package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"restfull_test/config"
	m "restfull_test/model"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
)

func TestInitUser(t *testing.T) {
	config.InitDB()
}

func TestGetUsersController(t *testing.T) {
	// config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/users", nil)
	response := httptest.NewRecorder()

	c := e.NewContext(request, response)

	if assert.NoError(t, GetUsersController(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
	}

}

func TestGetUserController(t *testing.T) {
	// config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("7")

	if assert.NoError(t, GetUserController(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
	}
}

func TestPostUserController(t *testing.T) {
	// config.InitDB()
	e := echo.New()

	mockUser := m.User{
		Name:     "Farhan",
		Email:    "farhan@gmail.com",
		Password: "1",
	}
	jsonEncode, _ := json.Marshal(mockUser)
	request := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(jsonEncode)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	err := CreateUserController(c)

	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, http.StatusCreated, response.Code)
	}

}

func TestUpdateUserController(t *testing.T) {
	// config.InitDB()
	e := echo.New()

	mockUser := m.User{
		Id:       7,
		Name:     "Farhan",
		Email:    "farhan@gmail.com",
		Password: "123",
	}
	jsonEncode, _ := json.Marshal(mockUser)
	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(jsonEncode)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("7")
	err := UpdateUserController(c)

	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, http.StatusOK, response.Code)
	}

}

func TestDeleteUserController(t *testing.T) {
	// config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodDelete, "/", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("9")

	if assert.NoError(t, DeleteUserController(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
	}
}

func TestLoginUserController(t *testing.T) {
	// config.InitDB()
	e := echo.New()
	userJSON := `{"email":"farhan@gmail.com","password":"123"}`
	request := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	err := LoginUserController(c)

	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, http.StatusOK, response.Code)
	}

}
