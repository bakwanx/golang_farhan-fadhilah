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

func TestInitBook(t *testing.T) {
	config.InitDB()
}

func TestGetBooksController(t *testing.T) {
	// config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/books", nil)
	response := httptest.NewRecorder()

	c := e.NewContext(request, response)

	if assert.NoError(t, GetBooksController(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
	}

}

func TestCreateBookController(t *testing.T) {
	// config.InitDB()
	e := echo.New()

	mockBook := m.Book{
		Title:  "Comic Naruto",
		Author: "Masashi Kishimoto",
	}
	jsonEncode, _ := json.Marshal(mockBook)
	request := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(string(jsonEncode)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	err := CreateBookController(c)

	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, http.StatusCreated, response.Code)
	}

}

func TestGetBookController(t *testing.T) {
	// config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")

	if assert.NoError(t, GetBookController(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
	}
}

func TestUpdateBookController(t *testing.T) {
	// config.InitDB()
	e := echo.New()

	mockBook := m.Book{
		Id:     2,
		Title:  "Manga Oshi Noko",
		Author: "Masashi Kishimoto",
	}
	jsonEncode, _ := json.Marshal(mockBook)
	request := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(jsonEncode)))
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	err := UpdateBookController(c)

	if err != nil {
		assert.Fail(t, err.Error())
	} else {
		assert.Equal(t, http.StatusOK, response.Code)
	}

}

func TestDeleteBookController(t *testing.T) {
	// config.InitDB()
	e := echo.New()
	request := httptest.NewRequest(http.MethodDelete, "/", nil)
	response := httptest.NewRecorder()
	c := e.NewContext(request, response)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("4")

	if assert.NoError(t, DeleteBookController(c)) {
		assert.Equal(t, http.StatusOK, response.Code)
	}
}
