package fetchers

import (
	"github.com/commune-project/commune/webfinger"
)

// ErrParsing describes that we cannot parse a subject.
var ErrParsing = webfinger.ErrParsing

// WebFinger combines two types of resources.
type WebFinger = webfinger.WebFinger

// QueryWebFinger fetches WebFinger response from remote.
func QueryWebFinger(subject string) (WebFinger, error) {
	return webfinger.QueryWebFinger(subject)
}

// ParseResponse returns a WebFinger struct
func ParseResponse(bytes []byte) (WebFinger, error) {
	return webfinger.ParseResponse(bytes)
}
