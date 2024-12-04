package routes

import (
	"tendanz/src/handlers"
	"tendanz/src/middleware"

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

	protected := g.Group("")
	protected.Use(middleware.ClientAuth)

	protected.POST("/verify", func(c echo.Context) error {
		return handlers.VerifyAccount(c, db)
	})

	protected.GET("/profile", func(c echo.Context) error {
		return handlers.GetProfile(c, db)	
	})

	return nil 
	
}