package routes

import (
	"tendanz/src/handlers"
	"tendanz/src/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AdminRoutes(g *echo.Group , db *gorm.DB) error {
	protectedAdmin := g.Group("")
	protectedAdmin.Use(middleware.AdminAuth)

	g.POST("/register", func(c echo.Context) error {
		return handlers.RegisterAdmin(c , db)
	})	
	g.POST("/login", func(c echo.Context) error {
		return handlers.LoginAdmin(c , db)
	})

	protectedAdmin.GET("/profile", func(c echo.Context) error {	
		return handlers.GetAdminProfile(c , db)
	})

	protectedAdmin.GET("/clients", func(c echo.Context) error {
		return handlers.GetAllClients(c , db)
	})

	protectedAdmin.GET("/stats" , func(c echo.Context) error {
		return handlers.AdminStat(c , db)	
	})


	return nil
}