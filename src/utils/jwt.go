package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GenerateToken(id uint) (string, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var secretKey string = os.Getenv("JWT_SECRET")
	if secretKey == "" {
		panic("Error getting the secret")
	}

	fmt.Println("The secret is:", secretKey)

	secretBytes := []byte(secretKey)

	claims := jwt.MapClaims{
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretBytes)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(token string) (jwt.MapClaims, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, fmt.Errorf("error getting the secret")
	}

	secretBytes := []byte(secretKey)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secretBytes, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("Token is not valid")
	}
	return claims, nil
}