package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

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
	if errors.Is(err, gorm.ErrRecordNotFound) {
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
