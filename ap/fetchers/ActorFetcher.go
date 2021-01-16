package fetchers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/commune-project/commune/ap/asparser"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/utils"
)

// ErrIsLocal describes that the uri is local.
var ErrIsLocal = errors.New("is local")

// GetOrFetchActorByURI gets a remote Account if uri is not found in the database.
func GetOrFetchActorByURI(context db.SiteContext, uri string) (interfaces.IActor, error) {
	account, err := dbmanagers.GetActorByURI(context, uri)
	// Already have the Account.
	if err == nil && account != nil {
		return account, nil
	}
	return FetchActorByURI(context, uri)
}

// FetchActorByURI gets a remote Account.
func FetchActorByURI(context db.SiteContext, uri string) (interfaces.IActor, error) {
	var domain string

	webfinger, err := QueryWebFinger(uri)
	if err != nil {
		u, err := url.Parse(uri)
		if err != nil {
			return nil, err
		}
		domain = u.Host
	} else {
		domain = webfinger.Domain
	}

	if utils.ContainsString(context.Settings.LocalDomains, domain) {
		return nil, ErrIsLocal
	}

	bytes, err := fetchIActorBytes(uri)
	if err != nil {
		return nil, err
	}

	var data interface{}
	err = json.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}

	if typedData, ok := data.(map[string]interface{}); ok {
		return asparser.ParseIActorWithDomain(typedData, domain), nil
	}

	return nil, ErrParsing
}

func fetchIActorBytes(uri string) ([]byte, error) {
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/activity+json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, int(resp.ContentLength))
	resp.Body.Read(bytes)
	return bytes, nil
}
