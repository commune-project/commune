package dbmanagers

import (
	"github.com/commune-project/commune/models"
	"gorm.io/gorm"
)

// GetCommuneByUsername returns the models.Commune named `username`.
func GetCommuneByUsername(db *gorm.DB, username string) (*models.Commune, error) {
	var commune models.Commune
	result := db.Where("username = ?", username).First(&commune)

	if result.Error != nil {
		return nil, result.Error
	}

	return &commune, nil
}
