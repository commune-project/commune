package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/commune-project/commune/ap/asgenerator"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/utils"
	"gorm.io/gorm"
)

type fetcher = func(r *http.Request) (interface{}, error)
type mapper = func(obj interface{}) map[string]interface{}

func writeError(w http.ResponseWriter, err error, errorCode int) {
	w.WriteHeader(errorCode)
	b, errjson := json.Marshal(map[string]string{
		"error": err.Error(),
	})
	if errjson != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(b)
	}
}

func apHandler(w http.ResponseWriter, r *http.Request, fetcher fetcher, mapper mapper) {
	obj, err := fetcher(r)
	if errors.Is(err, gorm.ErrRecordNotFound) || obj == nil {
		writeError(w, err, http.StatusNotFound)
		return
	} else if err != nil {
		writeError(w, err, http.StatusInternalServerError)
		return
	}
	mjson := mapper(obj)
	mjson["@context"] = "https://litepub.social/context.jsonld"
	b, err := json.Marshal(mjson)
	if err != nil {
		writeError(w, err, http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func genericMapper(obj interface{}) map[string]interface{} {
	return asgenerator.GenerateAS(obj)
}

func apOutboxHandler(url url.URL, actor interfaces.IActor) map[string]interface{} {
	query := url.Query()
	mjson := map[string]interface{}{
		"@context": "https://litepub.social/context.jsonld",
		"id":       actor.GetOutboxURI() + url.RawQuery,
	}

	if query.Get("page") != "true" {
		var totalItems int64 = 0
		if account, ok := actor.(*models.Account); ok {
			totalItems = db.DB.Model(account).Association("Posts").Count()
		} else if commune, ok := actor.(*models.Commune); ok {
			totalItems = db.DB.Model(commune).Association("Posts").Count()
		}
		mjson["type"] = "OrderedCollection"
		mjson["totalItems"] = totalItems
		mjson["first"] = actor.GetOutboxURI() + "?page=true"
		mjson["last"] = actor.GetOutboxURI() + "?min_id=0&page=true"
	} else {
		posts, _ := dbmanagers.GetPostsOfActor(db.DB.Scopes(utils.Paginate(query.Get("min_id"), query.Get("max_id"))).Order("id desc"), actor)
		mjson["type"] = "OrderedCollectionPage"
		mjson["partOf"] = actor.GetOutboxURI()
		if len(posts) > 0 {
			mjson["next"] = fmt.Sprintf("%s?max_id=%d&page=true", actor.GetOutboxURI(), posts[len(posts)-1].ID)
			mjson["prev"] = fmt.Sprintf("%s?min_id=%d&page=true", actor.GetOutboxURI(), posts[0].ID)
		}
		orderedItems := make([]interface{}, len(posts))
		for counter, post := range posts {
			orderedItems[counter] = asgenerator.GenerateAS(post.GetActivityCreate())
		}
		mjson["orderedItems"] = orderedItems
	}
	return mjson
}
