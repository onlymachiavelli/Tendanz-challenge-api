package services

import (
	"errors"
	"tendanz/src/models"

	"gorm.io/gorm"
)


type LifeInsuranceServiceIntr interface {


	GetOneLifeContract(id string, db *gorm.DB) (models.LifeInsurance, error)	
	CreateLifeContract(record models.LifeInsurance, db *gorm.DB) (models.LifeInsurance, error)
	GetAllLifeContracts(db *gorm.DB) ([]models.LifeInsurance, error)
	GetLifeContractByID(int, db *gorm.DB) (models.LifeInsurance, error)
	UpdateLifeContract(record models.LifeInsurance, db *gorm.DB) (models.LifeInsurance, error)
	DeleteLifeContract(int, db *gorm.DB) error
	AcceptLifeContract(int, db *gorm.DB) (models.LifeInsurance, error)
	RejectLifeContract(int, db *gorm.DB) (models.LifeInsurance, error)
	GetLifeContractByClientID(int, db *gorm.DB) ([]models.LifeInsurance, error)	

	GetLifeContractsByClient(clientID string, db *gorm.DB) ([]models.LifeInsurance, error)


}

type LifeInsuranceService struct {
}



func (u*LifeInsuranceService) GetOneLifeContract(id string, db *gorm.DB) (models.LifeInsurance, error) {
	if id == "" {
		return models.LifeInsurance{}, errors.New("id is required")
	}

	life := models.LifeInsurance{}
	errFinding := db.Where("id = ?", id).First(&life).Error

	if errFinding != nil {
		return models.LifeInsurance{}, errFinding
	}

	return life, nil

}


func (u*LifeInsuranceService) CreateLifeContract(record models.LifeInsurance, db *gorm.DB) (models.LifeInsurance, error) {
	errCreating := db.Create(&record).Error

	if errCreating != nil {
		return models.LifeInsurance{}, errCreating
	}

	return record, nil

}	

func (u *LifeInsuranceService)GetLifeContractsByClient(clientID string, db *gorm.DB) ([]models.LifeInsurance, error) {
	if clientID == "" {
		return []models.LifeInsurance{}, errors.New("client id is required")
	}

	lifeContracts := []models.LifeInsurance{}
	errFinding := db.Where("client_id = ?", clientID).Find(&lifeContracts).Error

	if errFinding != nil {
		return []models.LifeInsurance{}, errFinding
	}

	return lifeContracts, nil
}	

func (u* LifeInsuranceService)UpdateLifeContract(record models.LifeInsurance, db *gorm.DB) (models.LifeInsurance, error) {
	errUpdating := db.Save(&record).Error

	if errUpdating != nil {
		return models.LifeInsurance{}, errUpdating
	}

	return record, nil
}


func (u*LifeInsuranceService)DeleteLifeContract( record models.LifeInsurance, db *gorm.DB) error {
	errDeleting := db.Delete(&record).Error

	if errDeleting != nil {
		return errDeleting
 	}
	return nil
}