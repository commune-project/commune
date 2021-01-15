package dbmanagers

import (
	"github.com/commune-project/commune/models"
	"gorm.io/gorm"
)

// GetCommuneByUsername returns the models.Commune named `username`.
func GetCommuneByUsername(db *gorm.DB, username string, domain string) (*models.Commune, error) {
	var commune models.Commune
	result := db.Where("username = ? AND domain = ?", username, domain).First(&commune)

	if result.Error != nil {
		return nil, result.Error
	}

	return &commune, nil
}
