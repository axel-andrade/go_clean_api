package migrations

import (
	"go_clean_api/api/infra/database/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
}
