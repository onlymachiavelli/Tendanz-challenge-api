package main

import (
	"os"
	"tendanz/src/config"
	"tendanz/src/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func injectCors(e *echo.Echo) {
	devMode := true
	if (devMode) {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		  }))

	} else {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		  }))
	}
}

func main() {

	e := echo.New()

	injectCors(e)


	errLoading :=godotenv.Load() 
	if errLoading != nil {
		panic(errLoading)
	}
	PORT := os.Getenv("PORT")
	if PORT == "" { 
		panic("error loading the env")
	}

	//connect db 
	db, err := config.Connect()	
	if err != nil {	
		panic(err)
	}

	healthGroup := e.Group("/health")
	routes.HealthRoute(healthGroup)

	clientGroup := e.Group("/client")
	routes.ClientRoute(clientGroup, db)

	adminGroup := e.Group("/admin")	
	routes.AdminRoutes(adminGroup, db)




	e.Logger.Fatal(e.Start(":"+PORT))
}