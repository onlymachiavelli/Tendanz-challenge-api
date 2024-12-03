package services

import (
	"errors"
	"tendanz/src/models"

	"gorm.io/gorm"
)


type Service interface {	

	FindOneBy(field string , value string , db *gorm.DB) (models.Client,error)
	CreateRecord(db *gorm.DB , record models.Client)(models.Client , error)
	DeleteOne(db *gorm.DB, target string) (bool , error)

	//this one is only for test 

	UpdateOne(db *gorm.DB , record models.Client) (models.Client , error)
	GetAll(db *gorm.DB) ([]models.Client , error)	
}


type ServiceImpl struct{}

func (u *ServiceImpl) FindOneBy(field string , value string , db *gorm.DB) (models.Client , error) {
	if field == "" || value == "" {
		return models.Client{} , errors.New("please provide valid arguments")
	}
	target := models.Client{}
	errFinding := db.Where(field + " = ?" +value).First(&target).Error
	if (errFinding != nil) {
		return models.Client{} , errFinding
	}

	if target.ID == 0 {
		return models.Client{} , errors.New("record not found")
	}

	return  target , nil
}

func (u *ServiceImpl)CreateRecord(db *gorm.DB , record models.Client) ( models.Client , error) {

	if record.Email == "" || record.FirstName == "" || record.LastName =="" || record.Password =="" {
		return record,errors.New("please provide a valid payload") 
	}
	targetByPhone , errFindingByMail := u.FindOneBy("phone" , record.Email, db) 
	if errFindingByMail != nil {
		return models.Client{} ,errFindingByMail
	}


	if targetByPhone.ID != 0 {
		return models.Client{} , errors.New("phone already used")
	}

	targetByEmail , errFindingByMail := u.FindOneBy("email" , record.Email, db)

	if errFindingByMail != nil {
		return models.Client{} ,errFindingByMail
	}

	if targetByEmail.ID != 0 {
		return models.Client{} , errors.New("email already used")
	}	

	errCreating := db.Create(&record).Error	
	if errCreating != nil {	
		return models.Client{} , errCreating
	}

	return record , nil

}

func (u *ServiceImpl) DeleteRecord(db *gorm.DB , target string) (bool , error) {
	if target == "" {
		return false , errors.New("please provide a valid argument")
	}

	targetRecord , errFinding := u.FindOneBy("email" , target , db)	
	if errFinding != nil {	
		return false , errFinding
	}

	if targetRecord.ID == 0 {
		return false , errors.New("record not found")
	}

	errDeleting := db.Delete(&targetRecord).Error	
	if errDeleting != nil {
		return false , errDeleting
	}

	return true , nil
}

func (u *ServiceImpl)UpdateOne(db *gorm.DB , record models.Client) (models.Client , error) {
	if record.Email == "" {
		return models.Client{} , errors.New("please provide a valid argument")
	}

	targetRecord , errFinding := u.FindOneBy("email" , record.Email , db)	
	if errFinding != nil {
		return models.Client{} , errFinding
	}

	if targetRecord.ID == 0 {
		return models.Client{} , errors.New("record not found")
	}

	errUpdating := db.Model(&targetRecord).Updates(&record).Error	
	if errUpdating != nil {
		return models.Client{} , errUpdating
	}

	return targetRecord , nil
}

func GetAll(db *gorm.DB) ([]models.Client , error) {
	records := []models.Client{}
	errFinding := db.Find(&records).Error
	if errFinding != nil {
		return []models.Client{} , errFinding
	}

	return records , nil
}