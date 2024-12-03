package routes

import (
	"tendanz/src/handlers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ClientRoute(g *echo.Group, db *gorm.DB) error{


	g.POST("/register", func(c echo.Context) error {
		return handlers.Register(c, db)
	})

	g.POST("/login", func(c echo.Context) error {	
		return handlers.Login(c, db)
	})

	
	return nil 
	
	
}