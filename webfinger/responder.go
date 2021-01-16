package webfinger

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/commune-project/commune/models"
)

// Handler handles /.well-known/webfinger
func Handler(w http.ResponseWriter, r *http.Request) {
	var actor *models.Actor
	resource := r.URL.Query().Get("resource")
	acct, err := parseAcct(resource)
	if err == nil {
		actor, _ = dbmanagers.GetActorByUsername(db.Context.DB, acct.Username, acct.Domain)
	} else if strings.HasPrefix(resource, "https://") {
		actor, _ = dbmanagers.GetActorByURI(db.Context, resource)
	}

	if actor != nil {
		mjson := map[string]interface{}{
			"subject": "acct:" + actor.GetUsername() + "@" + actor.GetDomain(),
			"aliases": []string{
				actor.GetURI(),
				actor.GetURL(),
			},
			"links": []map[string]string{
				{
					"rel":  "http://webfinger.net/rel/profile-page",
					"type": "text/html",
					"href": actor.GetURL(),
				},
				{
					"rel":  "self",
					"type": "application/activity+json",
					"href": actor.GetURI(),
				},
			},
		}
		b, err := json.Marshal(mjson)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Add("Content-Type", "application/jrd+json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write(b)
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
