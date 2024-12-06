package handlers

import (
	"fmt"
	"tendanz/src/models"
	"tendanz/src/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

//by client
func GetStatsForClient(c echo.Context, db*gorm.DB) error {

	idClient := c.Get("client")	
	clientServices := services.ServiceImpl{}
	client,ErrFindingClient := clientServices.FindOneBy("id", fmt.Sprintf("%v" , idClient), db)	
	if ErrFindingClient != nil {	
		return c.JSON(400, map[string]interface{}{
			"message": "Client not found",
		})
	}

	lifes :=[] models.LifeInsurance{}

	errFinding := db.Where("client_id = ?", client.ID).Find(&lifes).Error	

	if errFinding != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "No contracts found",
		})
	}

	accepted := 0
	rejected := 0
	pending := 0
	total := 0

	for _, life := range lifes {
		if life.Status == "approved" {
			accepted++
		} else if life.Status == "rejected" {
			rejected++
		} else {
			pending++
		}
		total++
	}

	return c.JSON(200, map[string]interface{}{
		"accepted": accepted,
		"rejected": rejected,
		"pending": pending,
		"total": total,
	})

}	