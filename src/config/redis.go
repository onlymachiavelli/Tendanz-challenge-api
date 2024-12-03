package config

import "github.com/redis/go-redis/v9"

func ConnectRedis() (*redis.Client, error) {	

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", 
		DB:       0,  	
	})

	return rdb	, nil
}



