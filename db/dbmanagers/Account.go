package dbmanagers

import (
	"github.com/commune-project/commune/models"
	"gorm.io/gorm"
)

// GetAccountByUsername returns the models.Account named `username`.
func GetAccountByUsername(db *gorm.DB, username string) (*models.Account, error) {
	var account models.Account
	result := db.Where("username = ?", username).First(&account)

	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}
