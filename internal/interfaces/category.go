package interfaces

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sergiovirahonda/inventory-manager/internal/application"
	"gorm.io/gorm"
)

func Categories(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	categories, err := application.GetCategories(*user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, categories)
}

func GetCategoryById(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	categoryId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	category, err := application.GetCategoryById(categoryId, *user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusOK, category)
}

func CreateCategory(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	category := application.Category{}
	defer c.Request().Body.Close()
	err = json.NewDecoder(c.Request().Body).Decode(&category)
	if err != nil {
		log.Println("Failed reading the request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	instance, err := application.CreateCategory(category, *user)
	if err != nil {
		log.Println("Failed creating category:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusCreated, instance)
}

func DeleteCategoryById(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	categoryId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	err = application.DeleteCategoryById(categoryId, *user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	response := make(map[string]string)
	response["status"] = "deleted"
	return c.JSON(http.StatusOK, response)
}
