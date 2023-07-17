package controller

import (
	"clean_architecture/config"
	"clean_architecture/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestInitUser(t *testing.T) {
	config.ConnectDB()
}

func TestGetAllUsers(t *testing.T) {
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/users", nil)
	response := httptest.NewRecorder()

	c := e.NewContext(request, response)

	if assert.NoError(t, GetAllUsers(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
	}

}

func TestCreateUser(t *testing.T) {
	e := echo.New()

	mockUser := model.User{
		Email:    "test@gmail.com",
		Password: "123",
	}
	jsonEncode, _ := json.Marshal(mockUser)
	request := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(string(jsonEncode)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	err := CreateUser(c)

	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, http.StatusOK, response.Code)
	}

}

func TestLoginUserController(t *testing.T) {
	e := echo.New()
	userJSON := `{"email":"test@gmail.com","password":"123"}`
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
