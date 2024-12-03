package types

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	jwt.Claims
	UserID   int   `json:"user_id"`
	Email string `json:"email"`

}