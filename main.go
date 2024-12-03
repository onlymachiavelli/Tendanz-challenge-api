package main

import (
	"os"
	"tendanz/src/config"
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

	//connect db 
	db, err := config.Connect()	
	if err != nil {	
		panic(err)
	}

	healthGroup := e.Group("/health")
	routes.HealthRoute(healthGroup)

	clientGroup := e.Group("/client")
	routes.ClientRoute(clientGroup, db)



	e.Logger.Fatal(e.Start(":"+PORT))
}