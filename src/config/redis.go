package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func ConnectRedis() (*redis.Client, error) {	

	errLoading := godotenv.Load()
	if errLoading != nil {
		return nil, errLoading
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: "", 
		DB:       0,  	
	})

	return rdb	, nil
}



