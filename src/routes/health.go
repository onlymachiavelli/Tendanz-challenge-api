package routes

import (
	"tendanz/src/handlers"

	"github.com/labstack/echo/v4"
)

func HealthRoute(g *echo.Group) error {


	g.GET("" , func (c echo.Context) error  {
		return handlers.Healthy(c)
		
	})

	return nil 
}