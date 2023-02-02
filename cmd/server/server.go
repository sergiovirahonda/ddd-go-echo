package main

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/sergiovirahonda/inventory-manager/internal/domain/model"
	"github.com/sergiovirahonda/inventory-manager/internal/infrastructure"
	"github.com/sergiovirahonda/inventory-manager/internal/route"
)

func main() {
	// Echo instance
	e := route.Init()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// route.Restricted(e)
	db := infrastructure.NewDBInstance()
	db.SQL.AutoMigrate(&model.Users{})
	db.SQL.AutoMigrate(&model.Institutions{})
	db.SQL.AutoMigrate(&model.Categories{})
	db.SQL.AutoMigrate(&model.Services{})
	db.SQL.AutoMigrate(&model.Articles{})
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
