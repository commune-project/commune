package dbmanagers

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/commune-project/commune/ap"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"gorm.io/gorm"
)

// GetActorByID returns the models.Actor with `ID`.
func GetActorByID(db *gorm.DB, ID int) (*models.Actor, error) {
	var actor models.Actor
	result := db.First(&actor, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &actor, nil
}

// GetActorByUsername returns the models.Actor named `username`.
func GetActorByUsername(db *gorm.DB, username string, domain string) (*models.Actor, error) {
	var Actor models.Actor
	result := db.Where("lower(username) = lower(?) AND lower(domain) = lower(?)", username, domain).First(&Actor)

	if result.Error != nil {
		return nil, result.Error
	}

	return &Actor, nil
}

// GetActorByURI returns the models.Actor located at `uri`.
func GetActorByURI(context db.SiteContext, uri string) (actorP *models.Actor, err error) {
	if ap.IsLocal(context, uri) {
		myURL, err := url.Parse(uri)
		if err != nil {
			return nil, err
		}
		if re, err := regexp.Compile(fmt.Sprintf(`https://%s/[^/]+/([^/]+)`, myURL.Host)); err == nil {
			submatches := re.FindStringSubmatch(uri)
			if len(submatches) == 2 {
				return GetActorByUsername(context.DB, submatches[1], myURL.Host)
			}
		} else {
			return nil, err
		}
	} else {
		err = context.DB.Where("uri = ?", uri).First(actorP).Error
	}
	return
}
