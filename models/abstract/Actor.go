package abstract

import (
	"fmt"
)

// Actor is an abstract user either person, bot or a group of people. Don't create table for it.
type Actor struct {
	Object
	Username     string `json:"preferredUsername"`
	Domain       string `json:"-"`
	Name         string `json:"name"`
	Summary      string `json:"summary"`
	PublicKey    string `json:"-"`
	FollowersURI string `json:"followers"`
	FollowingURI string `json:"following"`
	InboxURI     string `json:"inbox"`
	OutboxURI    string `json:"outbox"`
}

func (actor *Actor) GetDomain() string {
	return actor.Domain
}

func (actor *Actor) GetURI() string {
	if actor.URI == nil {
		var slugType string
		if actor.Type == "Group" {
			slugType = "communities"
		} else {
			slugType = "users"
		}
		return fmt.Sprintf("https://%s/%s/%s", actor.Domain, slugType, actor.Username)
	} else {
		return *actor.URI
	}
}

func (actor *Actor) GetURL() string {
	if actor.URL == nil {
		return fmt.Sprintf("https://%s/@%s", actor.Domain, actor.Username)
	} else {
		return *actor.URL
	}
}

func (actor *Actor) GetUsername() string {
	return actor.Username
}

func (actor *Actor) RestChildren() map[string]interface{} {
	return map[string]interface{}{
		"name":    actor.Name,
		"summary": actor.Summary,
	}
}

func (actor *Actor) GetPublicKey() string {
	return actor.PublicKey
}

func (actor *Actor) GetPublicKeyURI() string {
	return actor.GetURI() + "#main-key"
}

func (actor *Actor) GetFollowersURI() string {
	if actor.FollowersURI != "" {
		return actor.FollowersURI
	}
	return actor.GetURI() + "/followers"
}

func (actor *Actor) GetFollowingURI() string {
	if actor.FollowingURI != "" {
		return actor.FollowingURI
	}
	return actor.GetURI() + "/following"
}

func (actor *Actor) GetInboxURI() string {
	if actor.InboxURI != "" {
		return actor.InboxURI
	}
	return actor.GetURI() + "/inbox"
}

func (actor *Actor) GetOutboxURI() string {
	if actor.OutboxURI != "" {
		return actor.OutboxURI
	}
	return actor.GetURI() + "/outbox"
}

func (actor *Actor) IsLocal(localDomains []string) bool {
	objDomain := actor.GetDomain()
	for _, localDomain := range localDomains {
		if objDomain == localDomain {
			return true
		}
	}
	return false
}

func (actor *Actor) IsBot() bool {
	return actor.Type == "Service"
}

func (actor *Actor) IsCommune() bool {
	return actor.Type == "Group"
}
