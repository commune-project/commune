package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/gorilla/mux"
)

// ApCommunityHandler handles AP requests to /communities/<username>
func ApCommunityHandler(w http.ResponseWriter, r *http.Request) {
	apHandler(w, r, getCommuneInterface, genericMapper)
}

func getCommuneInterface(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return dbmanagers.GetCommuneByUsername(db.DB(), vars["username"], r.Host)
}

// ApCommunityOutboxHandler handles AP requests to /communities/<username>/outbox
func ApCommunityOutboxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	commune, err := dbmanagers.GetCommuneByUsername(db.DB(), vars["username"], r.Host)

	if (err != nil) || (commune == nil) {
		writeError(w, err, http.StatusNotFound)
		return
	}

	b, err := json.Marshal(apOutboxHandler(*r.URL, commune))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/activity+json")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
