package dbmanagers

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/commune-project/commune/ap"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/models"
	"gorm.io/gorm"
)

// GetPostByID returns a certain post
func GetPostByID(db *gorm.DB, id int) (post models.Post, err error) {
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

// GetPostByURI returns the models.Post located at `uri` which is remote.
func GetPostByURI(context db.SiteContext, uri string) (*models.Post, error) {
	getLocalIDByURI := func(reStr string) (*models.Post, error) {
		re, err := regexp.Compile(reStr)
		if err != nil {
			return nil, err
		}

		submatches := re.FindStringSubmatch(uri)
		if len(submatches) == 2 {
			if id, err := strconv.Atoi(submatches[1]); err == nil {
				post, err := GetPostByID(context.DB, id)
				if err == nil {
					return &post, nil
				}
				return nil, err
			}
			return nil, err
		}
		return nil, errors.New("not found")
	}
	if ap.IsLocal(context, uri) {
		myURL, err := url.Parse(uri)
		if err != nil {
			return nil, err
		}
		if postP, err := getLocalIDByURI(fmt.Sprintf(`https://%s/p/([^/]+)`, myURL.Host)); err == nil {
			return postP, nil
		} else if postP, err := getLocalIDByURI(fmt.Sprintf(`https://%s/users/[^/]+/statuses/([^/]+)`, myURL.Host)); err == nil {
			return postP, nil
		} else {
			return nil, err
		}
	} else {
		var postP *models.Post
		err := context.DB.Where("uri = ?", uri).First(postP).Error
		return postP, err
	}
}
