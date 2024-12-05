package services

import (
	"tendanz/src/models"

	"gorm.io/gorm"
)

type AdminService struct{}

type AdminServiceInterface interface {

	CreateAdmin(record models.Admin , db*gorm.DB ) (models.Admin , error)
	UpdateAdmin(record models.Admin, db *gorm.DB) ( models.Admin , error)	
	DeleteAdmin(target string, db *gorm.DB) ( bool  , error)
	GetAllAdmins(db *gorm.DB) ( []models.Admin , error)
	FindAdminBy(field string, value string, db *gorm.DB) (models.Admin, error)	
}



func (u *AdminService) FindAdminBy(field string, value string, db *gorm.DB) ( models.Admin, error) {
	if field == "" || value == "" {
		return models.Admin{}, nil
	}
	target := models.Admin{}
	errFinding := db.Where(field+" = ?", value).First(&target).Error	
	return target ,  errFinding	
}

func (u *AdminService) CreateRecord(record models.Admin , db *gorm.DB) error {
	
	errCreating := db.Create(&record).Error
	return errCreating

}
