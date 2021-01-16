package dbmanagers

import (
	"errors"

	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/models"
	"gorm.io/gorm"
)

// GetPostByID returns a certain post
func GetPostByID(db *gorm.DB, id int64) (post models.Post, err error) {
	err = db.Model(&models.Post{}).Joins("Author").Joins("Category").Preload("Category.Commune").First(&post, id).Error
	return
}

// GetPostsOfAuthor returns all posts of the given Actor
func GetPostsOfAuthor(db *gorm.DB, actor models.Actor) ([]models.Post, error) {
	var posts []models.Post
	result := db.Where("author_id = ?", actor.ID).Preload("Author").Preload("Category").Preload("Category.Commune").Find(&posts)
	if result.Error != nil {
		return []models.Post{}, result.Error
	}
	return posts, nil
}

// GetPostsOfCommune returns all posts of the given Commune
func GetPostsOfCommune(db *gorm.DB, commune models.Actor) ([]models.Post, error) {
	var posts []models.Post
	result := db.Preload("Author").Joins("Category").Preload("Category.Commune").Find(&posts, "\"Category\".\"commune_id\" = ?", commune.ID)
	if result.Error != nil {
		return []models.Post{}, result.Error
	}
	return posts, nil
}

// GetPostsOfActor returns all posts of the given Actor
func GetPostsOfActor(db *gorm.DB, iActor interfaces.IActor) ([]models.Post, error) {
	actor, ok := iActor.(*models.Actor)
	if !ok {
		return nil, errors.New("??? not models.Actor?")
	}
	switch actor.GetType() {
	case "Group":
		return GetPostsOfCommune(db, *actor)
	case "Person":
		fallthrough
	case "Application":
		fallthrough
	case "Services":
		return GetPostsOfAuthor(db, *actor)
	default:
		return []models.Post{}, errors.New("forbidden")
	}
}
