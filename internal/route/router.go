package route

import (
	"github.com/labstack/echo/v4"
	"github.com/sergiovirahonda/inventory-manager/internal/config"
	"github.com/sergiovirahonda/inventory-manager/internal/interfaces"
)

func Init() *echo.Echo {
	e := echo.New()
	// Public routes
	e.GET("/api-auth/", interfaces.Login)
	// Users
	e.GET("/api/v0/users/", interfaces.Users, config.AuthValidation)
	e.GET("/api/v0/users/:id/", interfaces.GetUserById, config.AuthValidation)
	e.POST("/api/v0/users/", interfaces.CreateUser, config.AuthValidation)
	// Institution routes
	e.GET("/api/v0/institutions/", interfaces.Institutions, config.AuthValidation)
	e.GET("/api/v0/institutions/:id/", interfaces.GetInstitutionById, config.AuthValidation)
	e.POST("/api/v0/institutions/", interfaces.CreateInstitution, config.AuthValidation)
	// Category routes
	e.GET("/api/v0/categories/", interfaces.Categories, config.AuthValidation)
	e.GET("/api/v0/categories/:id/", interfaces.GetCategoryById, config.AuthValidation)
	e.POST("/api/v0/categories/", interfaces.CreateCategory, config.AuthValidation)
	e.DELETE("/api/v0/categories/", interfaces.DeleteCategoryById, config.AuthValidation)
	// Article routes
	e.GET("/api/v0/articles/", interfaces.Articles, config.AuthValidation)
	e.GET("/api/v0/articles/:id/", interfaces.GetArticleById, config.AuthValidation)
	e.GET("/api/v0/categories/:id/articles/", interfaces.GetArticlesFromCategory, config.AuthValidation)
	e.POST("/api/v0/articles/", interfaces.CreateArticle, config.AuthValidation)
	e.DELETE("/api/v0/articles/:id/", interfaces.DeleteArticleById, config.AuthValidation)
	// Service routes
	e.GET("/api/v0/services/", interfaces.Services, config.AuthValidation)
	e.GET("/api/v0/services/:id/", interfaces.GetServiceById, config.AuthValidation)
	e.GET("/api/v0/categories/:id/services/", interfaces.GetServicesFromCategory, config.AuthValidation)
	e.POST("/api/v0/services/", interfaces.CreateService, config.AuthValidation)
	e.DELETE("/api/v0/services/:id/", interfaces.DeleteServiceById, config.AuthValidation)
	return e
}
