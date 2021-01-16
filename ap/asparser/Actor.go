package asparser

import (
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/models/abstract"
	"github.com/commune-project/commune/utils"
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
	iActor := ParseIActor(data)
	if actor, ok := iActor.(*models.Actor); ok {
		actor.Domain = domain
		return actor
	}
	return iActor
}

func ParseIActor(data map[string]interface{}) interfaces.IActor {
	if dataType, ok := data["type"].(string); ok && utils.ContainsString(interfaces.ActorTypes, dataType) {
		return parseIntoActor(data)
	}
	return nil
}

func parseIntoActor(data map[string]interface{}) interfaces.IActor {
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

	return &models.Actor{
		Object: abstract.Object{
			Type: dataType,
			URI:  &dataURI,
			URL:  dataURL,
		},
		Username:     dataPreferredUsername,
		Name:         dataName,
		Summary:      dataSummary,
		PublicKey:    dataPublicKeyPEM,
		FollowersURI: dataFollowers,
		FollowingURI: dataFollowing,
		InboxURI:     dataInbox,
		OutboxURI:    dataOutbox,
	}
}
