package fetchers

import (
	"net/url"

	"github.com/commune-project/commune/ap/asparser"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/utils"
	"github.com/commune-project/commune/utils/commonerrors"
)

// GetOrFetchPostByURI gets a remote Post if uri is not found in the database.
func GetOrFetchPostByURI(context db.SiteContext, uri string) (interfaces.IPost, error) {
	post, err := dbmanagers.GetPostByURI(context, uri)
	// Already have the Account.
	if err == nil && post != nil {
		return post, nil
	}
	return FetchPostByURI(context, uri)
}

// FetchPostByURI gets a remote Post.
func FetchPostByURI(context db.SiteContext, uri string) (interfaces.IPost, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	if utils.ContainsString(context.Settings.LocalDomains, u.Host) {
		return nil, commonerrors.ErrIsLocal
	}

	if typedData, err := fetchJSON(uri); err == nil {
		return asparser.ParseIPost(typedData), nil
	}
	return nil, err
}
