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

func CreateLifeInsurranceContract(c echo.Context, db *gorm.DB) error {
	idClient := c.Get("client")
	if idClient == nil {
		return c.JSON(400, map[string]interface{}{
			"message": "client id is required",
		})
	}


	payload := requests.CREATELIFECONTRACT{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "invalid payload",
		})
	}


	lifeServices := services.LifeInsuranceService{}	
	clientServices := services.ServiceImpl{}

	client, err := clientServices.FindOneBy("id", fmt.Sprintf("%v" , idClient), db)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "client not found",
		})
	}

	lifeContract := models.LifeInsurance{
		ClientID: int(client.ID),
		PolicyType: payload.PolicyType,
		FaceAmount: payload.FaceAmount,
		PremiumMode: payload.PremiumMode,
		PremiumAmount: payload.PremiumAmount,
		PolicyTerm: payload.PolicyTerm,
		BenificiaryName: payload.BenificiaryName,
		BenificiaryRelationship: payload.BenificiaryRelationship,
		ContingentBenificiaryName: payload.ContingentBenificiaryName,
		ContingentBenificiaryRelationship: payload.ContingentBenificiaryRelationship,
		EffectiveDate: payload.EffectiveDate,
		ExpirationDate: payload.ExpirationDate,
		Status: "pending",
		Message: "",
		CreatedAt: time.Now(),
		UpdatedAt:time.Now(),		
	}

	createdLifeContract, err := lifeServices.CreateLifeContract(lifeContract, db)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error creating life contract",
		})
	}


	errSendingEmail := utils.ContractPendingMail(client.Email, "Life ", client.FirstName + " " + client.LastName)

	if errSendingEmail != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error sending email",
		})
	}
	
	return c.JSON(200, map[string]interface{}{
		"message": "life contract created",
		"data": createdLifeContract,
	})

}


func DeleteLifeInsurrance(c echo.Context, db *gorm.DB) error {
	return nil
}


func GetLifeInsurranceContract(c echo.Context, db *gorm.DB) error {
	return nil
}

func GetLifeInsurranceContracts(c echo.Context, db *gorm.DB) error {
	return nil
}


//as client 

func GetAllMyLifeContracts(c echo.Context, db *gorm.DB) error {
	idClient := c.Get("client")
	if idClient == nil {
		return c.JSON(400, map[string]interface{}{
			"message": "client id is required",
		})
	}

	lifeServices := services.LifeInsuranceService{}
	clientServices := services.ServiceImpl{}

	client, err := clientServices.FindOneBy("id", fmt.Sprintf("%v" , idClient), db)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "client not found",
		})
	}


	if client.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message": "client not found",
		})
	}

	lifeContracts, err := lifeServices.GetLifeContractsByClient(fmt.Sprintf("%v" , client.ID), db)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error getting life contracts",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "life contracts",
		"data": lifeContracts,
	})
}



//this is accessible by admin 


func GetLifeContractsForClient(c echo.Context , db*gorm.DB) error {

	return nil 
}


func AcceptLifeContract(c echo.Context, db *gorm.DB) error {
	adminId := c.Get("admin")
	if adminId == nil {	
		return c.JSON(
			400, map[string]interface{}{
				"message": "admin id is required",
			},
		)	
	}

	payload := requests.AcceptRejectLifeContract{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "invalid payload",
		})

	}


	adminServices := services.AdminService{}
	admin, err := adminServices.FindAdminBy("id", fmt.Sprintf("%v", adminId), db)

	if err != nil {	
		return c.JSON(
			400, map[string]interface{}{
				"message": "admin not found",
			},
		)		
	}

	if admin.ID == 0 {
		return c.JSON(
			400, map[string]interface{}{
				"message": "admin not found",
			},
		)

	}


	lifeServices := services.LifeInsuranceService{}
	lifeContractId := c.Param("id")

	if lifeContractId == "" {
		return c.JSON(400, map[string]interface{}{
			"message": "life contract id is required",
		})
	}

	lifeContract := models.LifeInsurance{}
	
	lifeContract, err = lifeServices.GetOneLifeContract(fmt.Sprintf("%v"  , lifeContractId) , db)

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "life contract not found",
		})
	}

	if lifeContract.ID == 0 {
		return c.JSON(400, map[string]interface{}{

			"message": "life contract not found",
		})
	}


	client := models.Client{}

	clientServices := services.ServiceImpl{}	
	
	client, err = clientServices.FindOneBy("id", fmt.Sprintf("%v" , lifeContract.ClientID), db)

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "client not found",
		})
	}

	if client.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message": "client not found",
		})
	}

	if lifeContract.Status == "approved" {
		return c.JSON(400, map[string]interface{}{
			"message": "life contract already approved",
		})
	}

	lifeContract.Status = "approved"
	lifeContract.Message = payload.Message

	updatedLifeContract, err := lifeServices.UpdateLifeContract(lifeContract, db)

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			
			"message": "error updating life contract",
		})
	}

	errSendingEmail := utils.ContractAccepted(client.Email, "Life ", client.FirstName + " " + client.LastName)

	if errSendingEmail != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error sending email",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "life contract approved",
		"data": updatedLifeContract,
	})

}	





func RejectLifeContract(c echo.Context, db *gorm.DB) error {
	adminId := c.Get("admin")
	if adminId == nil {	
		return c.JSON(
			400, map[string]interface{}{
				"message": "admin id is required",
			},
		)	
	}

	payload := requests.AcceptRejectLifeContract{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "invalid payload",
		})

	}


	adminServices := services.AdminService{}
	admin, err := adminServices.FindAdminBy("id", fmt.Sprintf("%v", adminId), db)

	if err != nil {	
		return c.JSON(
			400, map[string]interface{}{
				"message": "admin not found",
			},
		)		
	}

	if admin.ID == 0 {
		return c.JSON(
			400, map[string]interface{}{
				"message": "admin not found",
			},
		)

	}


	lifeServices := services.LifeInsuranceService{}
	lifeContractId := c.Param("id")

	if lifeContractId == "" {
		return c.JSON(400, map[string]interface{}{
			"message": "life contract id is required",
		})
	}

	lifeContract := models.LifeInsurance{}
	
	lifeContract, err = lifeServices.GetOneLifeContract(fmt.Sprintf("%v"  , lifeContractId) , db)

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "life contract not found",
		})
	}

	if lifeContract.ID == 0 {
		return c.JSON(400, map[string]interface{}{

			"message": "life contract not found",
		})
	}


	client := models.Client{}

	clientServices := services.ServiceImpl{}	
	
	client, err = clientServices.FindOneBy("id", fmt.Sprintf("%v" , lifeContract.ClientID), db)

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "client not found",
		})
	}

	if client.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message": "client not found",
		})
	}

	if lifeContract.Status == "approved" {
		return c.JSON(400, map[string]interface{}{
			"message": "life contract already approved",
		})
	}

	lifeContract.Status = "rejected"
	lifeContract.Message = payload.Message

	updatedLifeContract, err := lifeServices.UpdateLifeContract(lifeContract, db)

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			
			"message": "error updating life contract",
		})
	}

	errSendingEmail := utils.ContractRejected(client.Email, "Life ", client.FirstName + " " + client.LastName)

	if errSendingEmail != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error sending email",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "life contract rejected",
		"data": updatedLifeContract,
	})

}	