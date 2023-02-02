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

func Institutions(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	institutions, err := application.GetInstitutions(*user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, institutions)
}

func GetInstitutionById(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	institutionId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	institution, err := application.GetInstitutionById(institutionId, *user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusOK, institution)
}

func CreateInstitution(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	institution := application.Institution{}
	defer c.Request().Body.Close()
	err = json.NewDecoder(c.Request().Body).Decode(&institution)
	if err != nil {
		log.Fatalf("Failed reading the request body %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	instance, err := application.CreateInstitution(institution, *user)
	if err != nil {
		log.Println("Failed creating user:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusCreated, instance)
}
