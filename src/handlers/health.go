package handlers

import "github.com/labstack/echo/v4"

func Healthy(c echo.Context) error {
	return c.JSON(
		200 , 
		map[string]interface{}{
			"message" : "Iam Healthy" ,
		},
	)
}