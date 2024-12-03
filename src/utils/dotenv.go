package utils

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadVariable(key string) string {
	errLoadEnv := godotenv.Load()
	if errLoadEnv != nil {
		fmt.Errorf(errLoadEnv.Error())
		return ""
	}
	return os.Getenv(key)

}