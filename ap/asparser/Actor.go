package asparser

import (
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/models/abstract"
)

// Actor is an abstract user either person, bot or a group of people. Don't create table for it.
// type Actor struct {
// 	Object
// 	Username     string `json:"preferredUsername"`
// 	Domain       string `json:"-"`
// 	Name         string `json:"name"`
// 	Summary      string `json:"summary"`
// 	PublicKey    string `json:"-"`
// 	FollowersURI string `json:"followers"`
// 	FollowingURI string `json:"following"`
// 	InboxURI     string `json:"inbox"`
// 	OutboxURI    string `json:"outbox"`
// }

func ParseIActorWithDomain(data map[string]interface{}, domain string) interfaces.IActor {
	actor := ParseIActor(data)
	if account, ok := actor.(*models.Account); ok {
		account.Domain = domain
		return account
	} else if commune, ok := actor.(*models.Commune); ok {
		commune.Domain = domain
		return commune
	}
	return actor
}

func ParseIActor(data map[string]interface{}) interfaces.IActor {
	if dataType, ok := data["type"].(string); ok && dataType == "Group" {
		return parseIntoCommune(data)
	} else {
		return parseIntoAccount(data)
	}
}

func parseIntoAccount(data map[string]interface{}) interfaces.IActor {
	dataURI, _ := data["id"].(string)
	dataURL, _ := data["url"].(string)
	dataType, _ := data["type"].(string)
	dataPreferredUsername, _ := data["preferredUsername"].(string)
	dataName, _ := data["name"].(string)
	dataSummary, _ := data["summary"].(string)
	dataPublicKey, _ := data["publicKey"].(map[string]string)
	dataPublicKeyPEM, _ := dataPublicKey["publicKeyPem"]
	dataFollowers, _ := data["followers"].(string)
	dataFollowing, _ := data["following"].(string)
	dataInbox, _ := data["inbox"].(string)
	dataOutbox, _ := data["outbox"].(string)

	return &models.Account{
		Actor: abstract.Actor{
			Object: abstract.Object{
				Type: dataType,
				URI:  &dataURI,
				URL:  &dataURL,
			},
			Username:     dataPreferredUsername,
			Name:         dataName,
			Summary:      dataSummary,
			PublicKey:    dataPublicKeyPEM,
			FollowersURI: dataFollowers,
			FollowingURI: dataFollowing,
			InboxURI:     dataInbox,
			OutboxURI:    dataOutbox,
		},
	}
}

func parseIntoCommune(data map[string]interface{}) interfaces.IActor {
	dataURI, _ := data["id"].(string)
	dataURL, _ := data["url"].(string)
	dataPreferredUsername, _ := data["preferredUsername"].(string)
	dataName, _ := data["name"].(string)
	dataSummary, _ := data["summary"].(string)
	dataPublicKey, _ := data["publicKey"].(map[string]string)
	dataPublicKeyPEM, _ := dataPublicKey["publicKeyPem"]
	dataFollowers, _ := data["followers"].(string)
	dataFollowing, _ := data["following"].(string)
	dataInbox, _ := data["inbox"].(string)
	dataOutbox, _ := data["outbox"].(string)

	return &models.Commune{
		Actor: abstract.Actor{
			Object: abstract.Object{
				Type: "Group",
				URI:  &dataURI,
				URL:  &dataURL,
			},
			Username:     dataPreferredUsername,
			Name:         dataName,
			Summary:      dataSummary,
			PublicKey:    dataPublicKeyPEM,
			FollowersURI: dataFollowers,
			FollowingURI: dataFollowing,
			InboxURI:     dataInbox,
			OutboxURI:    dataOutbox,
		},
	}
}
