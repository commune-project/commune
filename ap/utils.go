package ap

import (
	"net/url"

	"github.com/commune-project/commune/db"
)

// IsLocal checks if a URI is of our server.
func IsLocal(context db.SiteContext, sURI string) bool {
	uri, err := url.Parse(sURI)
	if err != nil {
		return false
	}
	objDomain := uri.Host
	for _, localDomain := range context.Settings.LocalDomains {
		if objDomain == localDomain {
			return true
		}
	}
	return false
}

func SetContext(data map[string]interface{}) map[string]interface{} {
	data["@context"] = []interface{}{
		"https://www.w3.org/ns/activitystreams",
		"https://litepub.social/context.jsonld",
		map[string]string{
			"@language": "und",
		},
	}
	return data
}
