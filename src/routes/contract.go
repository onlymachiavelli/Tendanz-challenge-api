package routes

import (
	"tendanz/src/handlers"
	"tendanz/src/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ContractRoutes(e *echo.Group, db *gorm.DB) error{


	protectedClient := e.Group("")
	protectedClient.Use(middleware.ClientAuth)

	protectedAdmin := e.Group("")
	protectedAdmin.Use(middleware.AdminAuth)


	protectedClient.POST("/life", func(c echo.Context) error {
		return handlers.CreateLifeInsurranceContract(c, db)
	})

	protectedClient.GET("/life/mine", func (c echo.Context) error  {
		return handlers.GetLifeContractsAsClient(c , db)
		
	})
	protectedAdmin.GET("/life/client/:id" , func ( c echo.Context) error {
		return handlers.GetClientLifeInsurrance(c , db)
	})

	protectedAdmin.PUT("/life/:id/accept", func(c echo.Context) error {	
		return handlers.AcceptLifeContract(c, db)
	})

	protectedAdmin.PUT("/life/:id/reject", func(c echo.Context) error {	
		return handlers.RejectLifeContract(c, db)
	})

	return nil 
	
}