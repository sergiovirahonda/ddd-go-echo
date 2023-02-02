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

func Articles(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	articles, err := application.GetArticles(*user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	return c.JSON(http.StatusOK, articles)
}

func GetArticleById(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	articleId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	article, err := application.GetArticleById(articleId, *user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusOK, article)
}

func GetArticlesFromCategory(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	categoryId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	articles, err := application.GetArticlesFromCategory(categoryId, *user)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, err.Error)
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusOK, articles)
}

func CreateArticle(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	article := application.Article{}
	defer c.Request().Body.Close()
	err = json.NewDecoder(c.Request().Body).Decode(&article)
	if err != nil {
		log.Println("Failed reading the request body", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	instance, err := application.CreateArticle(article, *user)
	if err != nil {
		log.Println("Failed creating article:", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	return c.JSON(http.StatusCreated, instance)
}

func DeleteArticleById(c echo.Context) error {
	user, err := application.GetUserFromContext(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error)
	}
	id := c.Param("id")
	articleId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error)
	}
	err = application.DeleteArticleById(articleId, *user)
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
