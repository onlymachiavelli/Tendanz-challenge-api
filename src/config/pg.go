package config

import (
	"fmt"
	"os"
	"tendanz/src/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct { 
	Port string 
	Name string 
	Pass string 
	Host string 
	User string 
}

func Connect() (*gorm.DB, error) {

	errLoading := godotenv.Load()
	if errLoading != nil {
		return nil ,  errLoading
	}
	config := Config{
		Port: os.Getenv("PORTDB") ,
		Name:  os.Getenv("DBNAME"),
		Pass: os.Getenv("DBPASS"),
		Host:  os.Getenv("DBHOST"),
		User:  os.Getenv("DBUSER"),
	}

		if config.Port == "" || config.Name == "" || config.Pass == "" || config.Host == "" || config.User == "" {
			return nil, fmt.Errorf("missing environment	 variables")	
		}	

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s  sslmode=disable TimeZone=GMT+1",	
		config.Host, config.User, config.Pass, config.Name, config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	
	errMigratingClient := db.AutoMigrate(&models.Client{})
	if errMigratingClient != nil {
		return nil, errMigratingClient
	}

	return db, nil 
}

