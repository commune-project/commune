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

// GetCommuneByUsername returns the models.Commune named `username`.
func GetCommuneByUsername(db *gorm.DB, username string, domain string) (*models.Commune, error) {
	var commune models.Commune
	result := db.Where("username = ? AND domain = ?", username, domain).First(&commune)

	if result.Error != nil {
		return nil, result.Error
	}

	return &commune, nil
}

// GetCommuneByURI returns the models.Commune located at `uri`.
func GetCommuneByURI(context db.SiteContext, uri string) (accountP *models.Commune, err error) {
	if ap.IsLocal(context, uri) {
		myURL, err := url.Parse(uri)
		if err != nil {
			return nil, err
		}
		if re, err := regexp.Compile(fmt.Sprintf(`https://%s/communities/([^/]+)`, myURL.Host)); err == nil {
			submatches := re.FindStringSubmatch(uri)
			if len(submatches) == 2 {
				return GetCommuneByUsername(context.DB, submatches[1], myURL.Host)
			}
		} else {
			return nil, err
		}
	} else {
		err = context.DB.Where("uri = ?", uri).First(accountP).Error
	}
	return
}
