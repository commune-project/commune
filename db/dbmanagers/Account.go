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

// GetAccountByID returns the models.Account with `ID`.
func GetAccountByID(db *gorm.DB, ID int64) (*models.Account, error) {
	var account models.Account
	result := db.First(&account, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}

// GetAccountByUsername returns the models.Account named `username`.
func GetAccountByUsername(db *gorm.DB, username string, domain string) (*models.Account, error) {
	var account models.Account
	result := db.Where("username = ? AND domain = ?", username, domain).First(&account)

	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}

// GetAccountByURI returns the models.Account located at `uri`.
func GetAccountByURI(context db.SiteContext, uri string) (accountP *models.Account, err error) {
	if ap.IsLocal(context, uri) {
		myURL, err := url.Parse(uri)
		if err != nil {
			return nil, err
		}
		if re, err := regexp.Compile(fmt.Sprintf(`https://%s/users/([^/]+)`, myURL.Host)); err == nil {
			submatches := re.FindStringSubmatch(uri)
			if len(submatches) == 2 {
				return GetAccountByUsername(context.DB, submatches[1], myURL.Host)
			}
		} else {
			return nil, err
		}
	} else {
		err = context.DB.Where("uri = ?", uri).First(accountP).Error
	}
	return
}
