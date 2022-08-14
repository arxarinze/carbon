package migrations

import (
	"carbon/internal/models"

	"gorm.io/gorm"
)

func RunPasteMigration(db gorm.DB) {
	db.AutoMigrate(&models.Paste{})
}
