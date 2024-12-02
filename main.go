package main

import (
	"os"
	"tendanz/src/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()


	errLoading :=godotenv.Load() 
	if errLoading != nil {
		panic(errLoading)
	}
	PORT := os.Getenv("PORT")
	if PORT == "" { 
		panic("error loading the env")
	}

	//middlewares 
	healthGroup := e.Group("/health")
	routes.HealthRoute(healthGroup)


	e.Logger.Fatal(e.Start(":"+PORT))
}