package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/gorilla/mux"
)

// ApUserHandler handles AP requests to /users/<username>
func ApUserHandler(w http.ResponseWriter, r *http.Request) {
	apHandler(w, r, getAccountInterface, genericMapper)
}

func getAccountInterface(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	return dbmanagers.GetAccountByUsername(db.DB, vars["username"], r.Host)
}

// ApUserOutboxHandler handles AP requests to /users/<username>/outbox
func ApUserOutboxHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	account, err := dbmanagers.GetAccountByUsername(db.DB, vars["username"], r.Host)

	if (err != nil) || (account == nil) {
		writeError(w, err, http.StatusNotFound)
		return
	}

	b, err := json.Marshal(apOutboxHandler(*r.URL, account))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Add("Content-Type", "application/ld+json; profile=\"https://www.w3.org/ns/activitystreams\"")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
