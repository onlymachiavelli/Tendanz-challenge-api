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



//by client 
func DeleteLifeInsurrance(c echo.Context, db *gorm.DB) error {
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

	lifeContractId := c.Param("id")
	if lifeContractId == "" {
		return c.JSON(400, map[string]interface{}{
			"message": "life contract id is required",
		})
	}

	lifeContract, err := lifeServices.GetOneLifeContract(lifeContractId, db)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error getting life contract",
		})
	}

	if lifeContract.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message": "life contract not found",
		})
	}

	if lifeContract.ClientID != int(client.ID) {
		return c.JSON(400, map[string]interface{}{
			"message": "unauthorized",
		})
	}


	errDeleting := lifeServices.DeleteLifeContract(lifeContract, db)
	if errDeleting != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error deleting life contract",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "life contract deleted",
	})

}


func GetLifeInsurranceContract(c echo.Context, db *gorm.DB) error {
	return nil
}

func GetLifeInsurranceContracts(c echo.Context, db *gorm.DB) error {
	return nil
}


//as admin 

func GetClientLifeInsurrance(c echo.Context, db *gorm.DB) error {
	idAdmin := c.Get("admin")
	if idAdmin == nil {
		return c.JSON(400, map[string]interface{}{
			"message": "client id is required",
		})
	}

	adminServices := services.AdminService{}


	admin , errGettingAdmin := adminServices.FindAdminBy("id", fmt.Sprintf("%v", idAdmin) , db)
	
	if errGettingAdmin != nil {
		return c.JSON(
			400 , map[string]interface{}{
				"message" : "You're unauthorized, please login First",
			} ,
		)
	}

	if admin.ID == 0 {
		return c.JSON(
			404 , map[string]interface{}{
				"message" : "Admin Not Found",
			} ,
		)
	}

	idClient := c.Param("id")
	if idClient == "" {
		return c.JSON(
			400, 
			map[string]interface{}{
				"message" : "Please Select a Client",
			},
		)
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





//as client
func GetLifeContractsAsClient(c echo.Context , db*gorm.DB) error {

	idClient := c.Get("client")
	if idClient == nil {
		return c.JSON(400, map[string]interface{}{
			"message" : "Unauthorized, Please Login",
		})
	}

	clientServices := services.ServiceImpl{}


	client, errGettingClient := clientServices.FindOneBy("id",fmt.Sprintf("%v" , idClient), db) 

	if errGettingClient != nil  || client.ID == 0{
		return c.JSON(
			404, map[string]interface{}{
				"message":"Client is not found",
			},
		)
	}

	lifeContractServices := services.LifeInsuranceService{}

	lifes, errGettingLifes := lifeContractServices.GetLifeContractsByClient(fmt.Sprintf("%v", client.ID), db)
	
	if errGettingLifes != nil {
		return c.JSON(500 , map[string]interface{}{
			"message" : "Internal Error while getting your contracts, please report that !", 
		})
	}


	


	return c.JSON(
		200 , map[string]interface{}{
			"contracts" : lifes,
		},
	)

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




func GetOneLifeContractByClient(c echo.Context, db *gorm.DB) error {


	idClient := c.Get("client")	
	if idClient == nil {
		return c.JSON(400, map[string]interface{}{
			"message" : "Unauthorized, Please Login",
		})
	}

	clientServices := services.ServiceImpl{}

	client, errGettingClient := clientServices.FindOneBy("id",fmt.Sprintf("%v" , idClient), db)	

	if errGettingClient != nil  || client.ID == 0{
		return c.JSON(
			404, map[string]interface{}{
				"message":"Client is not found",
			},
		)
	}

	lifeContractServices := services.LifeInsuranceService{}

	lifeContractId := c.Param("id")
	if lifeContractId == "" {
		return c.JSON(400, map[string]interface{}{
			"message" : "Please select a contract",
		})
	}

	lifeContract, errGettingLifeContract := lifeContractServices.GetOneLifeContract(lifeContractId, db)

	if errGettingLifeContract != nil {
		return c.JSON(400, map[string]interface{}{
			"message" : "Error while getting the contract",
		})
	}

	if lifeContract.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message" : "Contract not found",
		})
	}

	if lifeContract.ClientID != int(client.ID) {
		return c.JSON(400, map[string]interface{}{
			"message" : "Unauthorized",
		})
	}

	return c.JSON(
		200 , map[string]interface{}{
			"contract" : lifeContract,
		},
	)


}




//client updates it 
func UpdateLifeContract(c echo.Context, db *gorm.DB) error {	

	idClient := c.Get("client")	


	//patch one
	payload := requests.UpdateLifeContract{}
	if idClient == nil {
		return c.JSON(400, map[string]interface{}{
			"message" : "Unauthorized, Please Login",
		})
	}


	clientServices := services.ServiceImpl{}

	client, errGettingClient := clientServices.FindOneBy("id",fmt.Sprintf("%v" , idClient), db)

	if errGettingClient != nil  || client.ID == 0{
		return c.JSON(
			404, map[string]interface{}{
				"message":"Client is not found",
			},
		)
	}

	lifeContractServices := services.LifeInsuranceService{}

	lifeContractId := c.Param("id")

	if lifeContractId == "" {
		return c.JSON(400, map[string]interface{}{
			"message" : "Please select a contract",
		})
	}



	lifeContract, errGettingLifeContract := lifeContractServices.GetOneLifeContract(lifeContractId, db)

	if errGettingLifeContract != nil {
		return c.JSON(400, map[string]interface{}{
			"message" : "Error while getting the contract",
		})
	}

	if lifeContract.ID == 0 {
		return c.JSON(400, map[string]interface{}{
			"message" : "Contract not found",
		})
	}

	if lifeContract.ClientID != int(client.ID) {
		return c.JSON(400, map[string]interface{}{
			"message" : "Unauthorized",
		})
	}


	if lifeContract.Status == "approved" {
		return c.JSON(400, map[string]interface{}{
			"message": "contract already approved",
		})
	}

	if err := c.Bind(&payload); err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "invalid payload",
		})
	}


	lifeContract.FaceAmount = payload.FaceAmount
	lifeContract.PremiumAmount = payload.PremiumAmount
	lifeContract.PolicyTerm = payload.PolicyTerm
	lifeContract.BenificiaryName = payload.BenificiaryName
	lifeContract.EffectiveDate = payload.EffectiveDate
	lifeContract.ExpirationDate = payload.ExpirationDate
	lifeContract.UpdatedAt = time.Now()
	updatedLifeContract, err := lifeContractServices.UpdateLifeContract(lifeContract, db)

	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error updating life contract",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"message": "life contract updated",
		"data": updatedLifeContract,
	})

}


func GetAllLifeContracts(c echo.Context, db *gorm.DB) error {
	idAdmin := c.Get("admin")	
	if idAdmin == nil {
		return c.JSON(400, map[string]interface{}{
			"message" : "Unauthorized, Please Login",

		})
	}

	adminServices := services.AdminService{}

	admin, errGettingAdmin := adminServices.FindAdminBy("id", fmt.Sprintf("%v", idAdmin) , db)	

	if errGettingAdmin != nil  || admin.ID == 0{
		return c.JSON(
			404, map[string]interface{}{
				"message":"Admin is not found",
			},
		)

	}

	lifeServices := services.LifeInsuranceService{}

	lifeContracts, err := lifeServices.GetAllLifeContracts(db)





	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"message": "error getting life contracts",
		})
	}


	data := []map[string]interface{}{}	

	for _, lifeContract := range lifeContracts {
		clientServices := services.ServiceImpl{}
		client, err := clientServices.FindOneBy("id", fmt.Sprintf("%v", lifeContract.ClientID), db)
		if err != nil {
			return c.JSON(400, map[string]interface{}{
				"message": "error getting client",
			})
		}

		data = append(data, map[string]interface{}{
			"client": client,
			"life_contract": lifeContract,
		})
	}



	return c.JSON(200, map[string]interface{}{
		"message": "life contracts",
		"data": data,
	})
}	