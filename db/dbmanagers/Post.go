package dbmanagers

import (
	"errors"

	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/models"
	"gorm.io/gorm"
)

// GetPostsOfAccount returns all posts of the given Account
func GetPostsOfAccount(db *gorm.DB, account models.Account) ([]models.Post, error) {
	var posts []models.Post
	result := db.Where("author_id = ?", account.ID).Preload("Author").Preload("Category").Preload("Category.Commune").Find(&posts)
	if result.Error != nil {
		return []models.Post{}, result.Error
	}
	return posts, nil
}

// GetPostsOfCommune returns all posts of the given Commune
func GetPostsOfCommune(db *gorm.DB, commune models.Commune) ([]models.Post, error) {
	var posts []models.Post
	result := db.Preload("Author").Joins("Category").Preload("Category.Commune").Find(&posts, "\"Category\".\"commune_id\" = ?", commune.ID)
	if result.Error != nil {
		return []models.Post{}, result.Error
	}
	return posts, nil
}

// GetPostsOfActor returns all posts of the given Actor
func GetPostsOfActor(db *gorm.DB, actor interfaces.IActor) ([]models.Post, error) {
	if account, ok := actor.(*models.Account); ok {
		return GetPostsOfAccount(db, *account)
	} else if commune, ok := actor.(*models.Commune); ok {
		return GetPostsOfCommune(db, *commune)
	} else {
		return []models.Post{}, errors.New("forbidden")
	}
}
