package migrators

import (
	"tendanz/src/models"

	"gorm.io/gorm"
)

func MigrateAdmin(db *gorm.DB) error {
	errMigrating := db.AutoMigrate(&models.Admin{})
	return errMigrating

}