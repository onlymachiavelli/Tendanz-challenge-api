package migrators

import (
	"tendanz/src/models"

	"gorm.io/gorm"
)
func MigrateUser(db *gorm.DB) error {

	errMigratingUser := db.AutoMigrate(&models.Client{})
	if errMigratingUser != nil {
		return errMigratingUser
	}

	return nil
}