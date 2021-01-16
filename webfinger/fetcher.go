package webfinger

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

// ErrParsing describes that we cannot parse a subject.
var ErrParsing = errors.New("error parsing")

type webFingerAcct struct {
	Username string
	Domain   string
}

type webFingerAP struct {
	URI string
	URL string
}

// WebFinger combines two types of resources.
type WebFinger struct {
	webFingerAP
	webFingerAcct
}

func (acct webFingerAcct) String() string {
	return "acct:" + acct.Username + "@" + acct.Domain
}

func (ap webFingerAP) String() string {
	return ap.URI
}

func parseAcct(subject string) (acct webFingerAcct, err error) {
	pair := strings.Split(subject[5:], "@")
	if len(pair) == 2 {
		acct = webFingerAcct{
			Username: pair[0],
			Domain:   pair[1],
		}
	} else {
		err = ErrParsing
	}
	return
}

const (
	typeAcct = iota
	typeAP   = iota
)

// QueryWebFinger fetches WebFinger response from remote.
func QueryWebFinger(subject string) (WebFinger, error) {
	var domain string
	if acct, err := parseAcct(subject); err == nil {
		domain = acct.Domain
	} else {
		u, err := url.Parse(subject)
		if err != nil {
			return WebFinger{}, nil
		}
		domain = u.Host
	}
	resp, err := http.Get("https://" + domain + "/.well-known/webfinger?resource=" + url.QueryEscape(domain))
	if err != nil {
		return WebFinger{}, nil
	}
	bytes := make([]byte, int(resp.ContentLength))
	resp.Body.Read(bytes)
	return ParseResponse(bytes)
}

// ParseResponse returns a WebFinger struct
func ParseResponse(bytes []byte) (webfinger WebFinger, err error) {
	response := webFingerResponse{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return
	}
	for _, link := range response.Links {
		if link.Type == "application/activity+json" {
			webfinger.URI = link.Href
		} else if link.Type == "text/html" {
			webfinger.URL = link.Href
		}
	}
	webfinger.webFingerAcct, err = parseAcct(response.Subject)
	return
}

type link struct {
	Rel  string `json:"rel"`
	Type string `json:"type"`
	Href string `json:"href"`
}

type webFingerResponse struct {
	Subject string `json:"subject"`
	Links   []link `json:"links"`
}
