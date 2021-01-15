package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/commune-project/commune/ap/asgenerator"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/commune-project/commune/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// ApPostHandler handles AP requests to /p/<id>
func ApPostHandler(w http.ResponseWriter, r *http.Request) {
	apHandler(w, r, getPost, genericMapper)
}

// ApPostActivityHandler handles AP requests to /p/<id>/activity
func ApPostActivityHandler(w http.ResponseWriter, r *http.Request) {
	apHandler(w, r, getPost, func(obj interface{}) map[string]interface{} {
		if value, ok := obj.(*models.Post); ok {
			return asgenerator.GenerateAS(value.GetActivityCreate())
		}
		return asgenerator.GenerateAS(obj)
	})
}

func getPost(r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	var post models.Post
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	post, err = dbmanagers.GetPostByID(db.DB, int64(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if vars["username"] != "" {
		if post.Author.GetUsername() != vars["username"] {
			return nil, errors.New("not found")
		}
	}

	return &post, nil
}
