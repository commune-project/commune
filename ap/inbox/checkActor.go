package inbox

import (
	"net/http"
	"net/url"

	"github.com/commune-project/commune/utils/commonerrors"
)

type checkActor struct{}

// CheckActor returns a new instance of checkActor.
func CheckActor() IInboxHandler {
	return checkActor{}
}

func (checkActor) Process(r *http.Request, data map[string]interface{}, processingInfo *ProcessingInfo) error {
	var actorURI string
	var dataURI string
	if dataID, ok := mapGetString(data, "id"); ok {
		dataURI = dataID
	} else {
		return commonerrors.ErrFormInvalid
	}
	if dataActorID, ok := data["actor"].(string); ok {
		actorURI = dataActorID
	} else if dataActorMap, ok := data["actor"].(map[string]interface{}); ok {
		if dataActorID, ok := dataActorMap["id"].(string); ok {
			actorURI = dataActorID
		}
	}
	if actorURI == "" {
		return commonerrors.ErrFormInvalid
	}

	if inSameDomain(dataURI, actorURI) {
		data["actor"] = actorURI
		return nil
	}
	return commonerrors.ErrCheckActorDataNotSameDomain
}

func inSameDomain(urlA string, urlB string) bool {
	if a, err := url.Parse(urlA); err == nil {
		if b, err := url.Parse(urlB); err == nil {
			if a.Host == b.Host {
				return true
			}
		}
	}
	return false
}
