package inbox

import (
	"errors"
	"net/http"
	"net/url"
)

var ErrCheckActorDataNoID = errors.New("data no id")
var ErrCheckActorDataNoActorID = errors.New("data no actor id")
var ErrCheckActorDataNotSameDomain = errors.New("data id and actor id is not on the same domain")

type checkActor struct{}

// CheckActor returns a new instance of checkActor.
func CheckActor() IInboxHandler {
	return checkActor{}
}

func (checkActor) Process(r *http.Request, data map[string]interface{}, object interface{}) error {
	var actorURI string
	var dataURI string
	if dataID, ok := mapGetString(data, "id"); ok {
		dataURI = dataID
	} else {
		return ErrCheckActorDataNoID
	}
	if dataActorID, ok := data["actor"].(string); ok {
		actorURI = dataActorID
	} else if dataActorMap, ok := data["actor"].(map[string]interface{}); ok {
		if dataActorID, ok := dataActorMap["id"].(string); ok {
			actorURI = dataActorID
		}
	}
	if actorURI == "" {
		return ErrCheckActorDataNoActorID
	}

	if urlID, err := url.Parse(dataURI); err == nil {
		if urlActorID, err := url.Parse(actorURI); err == nil {
			if urlActorID.Host == urlID.Host {
				return nil
			}
		}
	}
	return ErrCheckActorDataNotSameDomain
}
