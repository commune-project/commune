package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/commune-project/commune/ap/asgenerator"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// ApPostHandler handles AP requests to /p/<id>
func ApPostHandler(w http.ResponseWriter, r *http.Request) {
	apPostHandler(w, r, func(obj *models.Post) map[string]interface{} {
		return asgenerator.GenerateAS(obj)
	})
}

// ApPostActivityHandler handles AP requests to /p/<id>/activity
func ApPostActivityHandler(w http.ResponseWriter, r *http.Request) {
	apPostHandler(w, r, func(obj *models.Post) map[string]interface{} {
		return asgenerator.GenerateAS(obj.GetActivityCreate())
	})
}

func apPostHandler(w http.ResponseWriter, r *http.Request, f func(obj *models.Post) map[string]interface{}) {
	vars := mux.Vars(r)
	var post models.Post
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Record not found"))
		w.Write([]byte(err.Error()))
		return
	}
	result := db.DB.Model(&models.Post{}).Joins("Author").Joins("Category").First(&post, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Record not found"))
		w.Write([]byte(result.Error.Error()))
		return
	}

	mjson := f(&post)
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
