package middleware

import (
	"tendanz/src/utils"

	"github.com/labstack/echo/v4"
)

func AdminAuth(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {

		token := c.Request().Header.Get("Authorization")

		if token == "" {
			return c.JSON(400, map[string]interface{}{
				"message": "Token is required",
			})
		}

		token = token[7:]

		claims, err := utils.VerifyToken(token)
		if err != nil {
			return c.JSON(400, map[string]interface{}{
				"message": "Invalid token",
			})
		}
		c.Set("admin", int(claims["id"].(float64)))
		return next(c)

	}
}
