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

func Services(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	services, err := application.GetServices(*user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, services)
}

func GetServiceById(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	serviceId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	service, err := application.GetServiceById(serviceId, *user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusOK, service)
}

func GetServicesFromCategory(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	categoryId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	services, err := application.GetServicesFromCategory(categoryId, *user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusOK, services)
}

func CreateService(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	service := application.Service{}
	defer c.Request().Body.Close()
	err = json.NewDecoder(c.Request().Body).Decode(&service)
	if err != nil {
		log.Println("Failed reading the request body:", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	instance, err := application.CreateService(service, *user)
	if err != nil {
		log.Println("Failed creating service:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusCreated, instance)
}

func DeleteServiceById(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	serviceId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	err = application.DeleteServiceById(serviceId, *user)
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
