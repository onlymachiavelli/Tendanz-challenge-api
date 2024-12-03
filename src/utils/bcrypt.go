package utils

import (
	"tendanz/src/types"

	"fmt"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"

	"github.com/joho/godotenv"
)

func HashPassword(password string) (string , error) {

	loadEnvErr := godotenv.Load()
	if (loadEnvErr != nil) {
		fmt.Println("Error Loadind the local env")
		panic(loadEnvErr)
	}

	bKey := os.Getenv("SALT_ROUNDS")
	if bKey == "" {
		panic("Error Loading the BCRYPT Key")
	}
	key, err := strconv.Atoi(bKey)

	if (err != nil) {
		panic(err)
	}
	hashConfig := &types.Cost{
		Cost: key,
	}
	
	passwordWithSecret := append([]byte(password),byte(hashConfig.Cost))
	
	hash, err := bcrypt.GenerateFromPassword(passwordWithSecret, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil

}

func Verify(password, hashedPassword string) bool{
	loadEnvErr := godotenv.Load()
	if (loadEnvErr != nil) {
		fmt.Println("Error Loadind the local env")
		panic(loadEnvErr)
	}

	bKey := os.Getenv("SALT_ROUNDS")
	if bKey == "" {
		panic("Error Loading the BCRYPT Key")	
	}
	key, err := strconv.Atoi(bKey)

	if (err != nil) {
		panic(err)
	}
	hashConfig := &types.Cost{
		Cost: key,
	}
	inputPasswordWithSecret := append([]byte(password), byte(hashConfig.Cost))
	
	errPass := bcrypt.CompareHashAndPassword([]byte(hashedPassword), inputPasswordWithSecret)
	return errPass == nil
}