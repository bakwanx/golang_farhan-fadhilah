package controller

import (
	"net/http"
	"restfull_test/config"
	"restfull_test/model"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success all books",
		"data":    books,
	})
}

func GetBookController(c echo.Context) error {
	id := c.Param("id")
	var book model.Book

	if err := config.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"book":    book,
	})
}

func CreateBookController(c echo.Context) error {

	book := model.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	var book model.Book

	if err := config.DB.Delete(&book, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete",
	})
}

func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	var book model.Book

	if err := config.DB.Where("id = ?", id).First(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update book by id",
		"book":    book,
	})
}
