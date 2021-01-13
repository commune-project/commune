package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/commune-project/commune/ap/asgenerator"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ApUserHandler handles AP requests to /users/<username>
func ApUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var account models.Account
	result := db.DB.Model(&models.Account{}).Where("username = ?", vars["username"]).First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(result.Error.Error()))
		return
	}

	mjson := asgenerator.GenerateAS(&account)
	mjson["@context"] = "https://litepub.social/context.jsonld"
	b, err := json.Marshal(mjson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Add("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

// ApUserOutboxHandler handles AP requests to /users/<username>/outbox
func ApUserOutboxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var account models.Account
	result := db.DB.Model(&models.Account{}).Where("username = ?", vars["username"]).First(&account)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	query := r.URL.Query()

	mjson := map[string]interface{}{
		"@context": "https://litepub.social/context.jsonld",
		"id":       account.GetOutboxURI() + r.URL.RawQuery,
	}
	if query.Get("page") != "true" {
		mjson["type"] = "OrderedCollection"
		mjson["totalItems"] = db.DB.Model(&account).Association("Posts").Count()
		mjson["first"] = account.GetOutboxURI() + "?page=true"
		mjson["last"] = account.GetOutboxURI() + "?min_id=0&page=true"
	} else {
		var posts []models.Post
		db.DB.Scopes(utils.Paginate(query.Get("min_id"), query.Get("max_id"))).Order("ID desc").Preload(clause.Associations).Find(&posts)
		mjson["type"] = "OrderedCollectionPage"
		mjson["partOf"] = account.GetOutboxURI()
		if len(posts) > 0 {
			mjson["next"] = fmt.Sprintf("%s?max_id=%d&page=true", account.GetOutboxURI(), posts[len(posts)-1].ID)
			mjson["prev"] = fmt.Sprintf("%s?min_id=%d&page=true", account.GetOutboxURI(), posts[0].ID)
		}
		orderedItems := make([]interface{}, len(posts))
		for counter, post := range posts {
			orderedItems[counter] = asgenerator.GenerateAS(post.GetActivityCreate())
		}
		mjson["orderedItems"] = orderedItems
	}
	b, err := json.Marshal(mjson)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
