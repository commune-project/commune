package utils

import (
	"strconv"

	"gorm.io/gorm"
)

// Paginate any Model.
// db.Scopes(Paginate(r)).Find(&users)
// db.Scopes(Paginate(r)).Find(&articles)
func Paginate(sMinID string, sMaxID string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		result := db
		minID, err := strconv.Atoi(sMinID)
		if err == nil {
			result = result.Where("id > ?", minID)
		}
		maxID, err := strconv.Atoi(sMaxID)
		if err == nil {
			result = result.Where("id < ?", maxID)
		}
		return result.Limit(20)
	}
}
