package handlers

import (
	"fmt"
	"tendanz/src/config"
	"tendanz/src/models"
	"tendanz/src/services"
	"tendanz/src/types/requests"
	"tendanz/src/utils"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(c echo.Context, db *gorm.DB) error {

	fmt.Println("registering client")
	payload := requests.ClientRegister{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "invalid payload",
		})
	}
	
	if payload.Email == "" || payload.FirstName == "" || payload.LastName == "" || payload.Password == "" {
		return c.JSON(400, map[string]interface{}{
			"message": "please provide a valid payload",
		})
	}

	recordRow := models.Client{
		Email: payload.Email,
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Phone: payload.Phone,
		Birthdate: payload.Birthdate,
		Verified: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	hashedPass, errHashingPass := utils.HashPassword(payload.Password)
	if errHashingPass != nil {	
		return c.JSON(400, map[string]interface{}{
			"message": errHashingPass.Error(),
		})
	}
	recordRow.Password = hashedPass

	clientServices := services.ServiceImpl{}
	_, err := clientServices.CreateRecord(db, recordRow)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	code := utils.GenerateCode()
	if code == "" {
		return c.JSON(400, map[string]interface{}{
			"message": "error generating code",
		})
	}

	rds, errConnecting := config.ConnectRedis()
	if errConnecting != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error connecting to redis",
		})
	}

	errSetting := rds.Set(
		c.Request().Context(),
		recordRow.Email,
		code,
		time.Minute * 5,
	).Err()
	if errSetting != nil {
		fmt.Println("code issue : ",errSetting)	
		return c.JSON(400, map[string]interface{}{

			"message": "error setting code",
		})
	}


	errSending := utils.SendCode(recordRow.Email,  code)
	if errSending != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error sending email",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "record created successfully",
	})

}


func Login(c echo.Context, db *gorm.DB) error {

	payload := requests.LoginClient{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "invalid payload",
		})
	}

	if payload.Email == "" || payload.Password == "" {
		return c.JSON(400, map[string]interface{}{
			"message": "please provide a valid payload",
		})
	}

	clientServices := services.ServiceImpl{}
	target, err := clientServices.FindOneBy("email", payload.Email, db)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if target.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message": "record not found",
		})
	}

	if !utils.Verify(payload.Password, target.Password) {
		return c.JSON(400, map[string]interface{}{
			"message": "invalid credentials",
		})
	}

	token, err := utils.GenerateToken(target.ID)	
	if err != nil {
		return c.JSON(400, map[string]interface{}{

			"message": "error generating token",
		})
	}

	user := map[string]interface{}{
		"id": target.ID,
		"email": target.Email,
		"first_name": target.FirstName,
		"last_name": target.LastName,
		"phone": target.Phone,
		"verified": target.Verified,
		"created_at": target.CreatedAt,
		"updated_at": target.UpdatedAt,
	}
	return c.JSON(200, map[string]interface{}{
		"message": "login successful",
		"token": token,
		"user": user,
	})
}

func Auth(c echo.Context, db *gorm.DB) error {
	return nil
}	

func Update(c echo.Context , db *gorm.DB) error {
	return nil
}


func Delete(c echo.Context , db *gorm.DB) error {
	return nil
}

func VerifyAccount(c echo.Context , db *gorm.DB) error {

	payload := requests.VerifyCode{}	


	if err := c.Bind(&payload); err != nil {	
		return c.JSON(400, map[string]interface{}{
			"message": "invalid payload",
		})
	}

	if payload.Code == "" {
		return c.JSON(
			400,
			map[string]interface{}{
				"message": "please provide a valid code",	
			})
	}

	idClient := c.Get("client")
	if idClient == nil {
		return c.JSON(
			401, 
			map[string]interface{}{
				"message": "unauthorized",
			},
		)
	}

	clientServices := services.ServiceImpl{}	
	target, errFinding := clientServices.FindOneBy("id", fmt.Sprintf("%v" , idClient), db)	
	if errFinding != nil {
		return c.JSON(400, map[string]interface{}{
			"message": errFinding.Error(),
		})
	}

	if target.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message": "record not found",
		})	
	}

	rds, errConnecting := config.ConnectRedis()	
	if errConnecting != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error connecting to redis",
		})
	}

	code, errGetting := rds.Get(c.Request().Context(), target.Email).Result()		
	if errGetting != nil || code == "" {
		return c.JSON(400, map[string]interface{}{
			"message": "Code expired",
		})
	}


	if code != payload.Code {
		return c.JSON(400, map[string]interface{}{
			"message": "invalid code",
		})
	}	


	errDeleting := rds.Del(c.Request().Context(), target.Email).Err()	
	if errDeleting != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error deleting code",
		})
	}

	target.Verified = true	
	//change it later 

	errUpdating := db.Save(&target).Error
	if errUpdating != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error updating record",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "account verified",
	})
}

func GetProfile(c echo.Context , db *gorm.DB) error {

	idClient := c.Get("client")	

	if idClient == nil {
	
		return c.JSON(401, map[string]interface{}{
			"message": "unauthorized",
		})

	}

	clientServices := services.ServiceImpl{}	
	target, errFinding := clientServices.FindOneBy("id", fmt.Sprintf("%v" , idClient), db)
	if errFinding != nil {
		return c.JSON(400, map[string]interface{}{
			"message": errFinding.Error(),
		})
	}

	if target.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message": "record not found",
		})
	}

	data := map[string]interface{}{	
		"id": target.ID,
		"email": target.Email,
		"first_name": target.FirstName,
		"last_name": target.LastName,
		"phone": target.Phone,
		"verified": target.Verified,
		"created_at": target.CreatedAt,
		"updated_at": target.UpdatedAt,
	}

	return c.JSON(200, data)

}	

func ResendVerificationCode( c echo.Context , db*gorm.DB) error {
	idClient := c.Get("client")
	if idClient == nil {
		return c.JSON(
			401,
			map[string]interface{}{
				"message" : "Not Authorized !" ,
			},
		)
	}


	clientServices := services.ServiceImpl{}	


	target, errFinding := clientServices.FindOneBy("id", fmt.Sprintf("%v" , idClient), db)	

	if errFinding != nil {
		return c.JSON(400, map[string]interface{}{
			"message": errFinding.Error(),
		})

	}

	if target.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message": "Client not Found",
		})
	}

	code := utils.GenerateCode()	
	if code == "" {	
		return c.JSON(400, map[string]interface{}{
			"message": "error generating code",
		})
	}


	rds, errConnecting := config.ConnectRedis()
	if errConnecting != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error connecting to redis",
		})
	}

	errSetting := rds.Set(
		c.Request().Context(),
		target.Email,
		code,
		time.Minute * 5,
	).Err()
	if errSetting != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error setting code",
		})
	}

	errSending := utils.SendCode(target.Email,  code)
	if errSending != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error sending email",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "code sent successfully",
	})

}

func GetAllClients(c echo.Context , db *gorm.DB) error {
	idAdmin := c.Get("admin")	
	if idAdmin == nil {
		return c.JSON(401, map[string]interface{}{
			"message": "unauthorized",
		})
	}

	clientServices := services.ServiceImpl{}

	clients, err := clientServices.GetAll(db)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(200, clients)

}	