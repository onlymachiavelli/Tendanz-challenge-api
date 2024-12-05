package handlers

import (
	"fmt"
	"tendanz/src/models"
	"tendanz/src/services"
	"tendanz/src/types/requests"
	"tendanz/src/utils"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterAdmin(c echo.Context, db *gorm.DB) error {
	payload := requests.REGISTERADMINREQUEST{}	
	if err := c.Bind(&payload); err != nil {	
		return c.JSON(400, 
			map[string]interface{}{
				"message": "Invalid payload",
				"error": err,
			})
	}

	if payload.Email == "" || payload.Password == "" || payload.Identity == "" || payload.FirstName == "" || payload.LastName == "" || payload.PhoneNumber == "" {	
		return c.JSON(400,
			map[string]interface{}{
				"message": "Missing required fields",
			})
	}

	adminServices := services.AdminService{}	
	record := models.Admin{
		Email: payload.Email,
		Identity: payload.Identity,
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Phone: payload.PhoneNumber,
		Verified: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),	

	}
	hashed, errHashing := utils.HashPassword(payload.Password)	
	if errHashing != nil {	
		return c.JSON(500,
			map[string]interface{}{
				"message": "Error hashing password",
				"error": errHashing,
			})
	}
	record.Password = hashed

	errCreating := adminServices.CreateRecord(record, db)	
	if errCreating != nil {
		return c.JSON(500,
			map[string]interface{}{
				"message": "Error creating record",
				"error": errCreating,
			})
	}

	return c.JSON(200,
		map[string]interface{}{
			"message": "Admin created successfully",
			"record": record,
		})

}


func LoginAdmin(c echo.Context,db*gorm.DB) error {

	payload := requests.LOGINADMIN{}	
	
	if err := c.Bind(&payload); err != nil {	
		return c.JSON(400,
			map[string]interface{}{
				"message": "Invalid payload",
				"error": err,
			})
	}

	if payload.Email == "" || payload.Password == "" {
		return c.JSON(400,
			map[string]interface{}{
				"message": "Missing required fields",
			})
	}

	adminServices := services.AdminService{}

	record, errFinding := adminServices.FindAdminBy("email", payload.Email, db)
	if errFinding != nil {
		return c.JSON(500,
			map[string]interface{}{
				"message": "Error finding record",
				"error": errFinding,
			})
	}

	if record.ID == 0 {
		return c.JSON(404,
			map[string]interface{}{
				"message": "Record not found",
			})
	}

	if !utils.Verify(payload.Password, record.Password) {
		return c.JSON(401,
			map[string]interface{}{
				"message": "Invalid credentials",
			})
	}

	admin := map[string]interface{}{	
		"id": record.ID,
		"email": record.Email,
		"identity": record.Identity,
		"first_name": record.FirstName,
		"last_name": record.LastName,
		"phone": record.Phone,
		"verified": record.Verified,
		"created_at": record.CreatedAt,
		"updated_at": record.UpdatedAt,
	}

	token, errToken := utils.GenerateToken(record.ID)
	if errToken != nil {
		return c.JSON(500,
			map[string]interface{}{
				"message": "Error generating token",
				"error": errToken,
			})
	}
	return c.JSON(200,
		map[string]interface{}{
			"message": "Login successful",
			"admin": admin,
			"token": token,
		})
	
}

func GetAdminProfile(c echo.Context, db *gorm.DB) error {

	adminServices := services.AdminService{}
	idAdmin := c.Get("admin")
	if idAdmin == nil {
		return c.JSON(401,
			map[string]interface{}{
				"message": "Unauthorized",
			})
	}

	record, errFinding := adminServices.FindAdminBy("id", fmt.Sprintf("%v" , idAdmin), db)	
	if errFinding != nil {
		return c.JSON(500,
			map[string]interface{}{
				"message": "Error finding record",
				"error": errFinding,
			})
	}

	if record.ID == 0 {
		return c.JSON(404,
			map[string]interface{}{
				"message": "Record not found",
			})
	}

	admin := map[string]interface{}{
		"id": record.ID,
		"email": record.Email,
		"identity": record.Identity,
		"first_name": record.FirstName,
		"last_name": record.LastName,
		"phone": record.Phone,
		"verified": record.Verified,
		"created_at": record.CreatedAt,
		"updated_at": record.UpdatedAt,
	}
	
	return c.JSON(200,
		map[string]interface{}{
			"message": "Profile found",
			"admin": admin,
			})	
}