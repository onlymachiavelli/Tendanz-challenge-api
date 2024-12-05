package migrators

import (
	"tendanz/src/models"

	"gorm.io/gorm"
)

func MigrateContracts(db *gorm.DB)error {
	errMigratingLifeInsurance := db.AutoMigrate(&models.LifeInsurance{})

	if errMigratingLifeInsurance != nil {

		return errMigratingLifeInsurance
	}

	return nil 

}