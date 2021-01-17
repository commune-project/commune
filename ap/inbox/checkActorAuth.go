package inbox

import (
	"net/http"

	"github.com/commune-project/commune/ap/fetchers"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/utils/commonerrors"
)

type checkActorAuth struct {
	context db.SiteContext
}

// CheckActorAuth returns a new instance of checkActorAuth.
func CheckActorAuth(context db.SiteContext) IInboxHandler {
	return checkActorAuth{
		context: context,
	}
}

func (auth checkActorAuth) Process(r *http.Request, data map[string]interface{}, processingInfo *ProcessingInfo) error {
	if actorID, ok := mapGetString(data, "actor"); ok {
		iActor, err := fetchers.GetOrFetchActorByURI(auth.context, actorID)
		if err != nil {
			return err
		}
		if actor, ok := iActor.(*models.Actor); ok {
			processingInfo.Actor = actor
		}
	}
	return commonerrors.ErrNotLoggedIn
}
