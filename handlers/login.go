package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/utils"
	"github.com/commune-project/commune/utils/commonerrors"
)

type loginForm struct {
	// Can be a username or an email address.
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login reads from POSTed JSON and authenticate.
func Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if r.Header.Get("Content-Type") == "application/json" {
			form := loginForm{}
			decoder := json.NewDecoder(r.Body)
			if err := decoder.Decode(&form); err != nil {
				writeError(w, commonerrors.ErrFormInvalid, http.StatusBadRequest)
				log.Println(err)
				return
			}
			user, err := getUserByEmail(form.Username)
			if err != nil || user == nil {
				if user, err = getUserByUsername(form.Username, r.URL.Host); err != nil || user == nil {
					writeError(w, commonerrors.ErrAuthenticationFailed, http.StatusForbidden)
					return
				}
			}
			if user != nil {
				if err := user.Authenticate(form.Password); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				} else {
					session, err := db.Context.Store.Get(r, "web")
					if err != nil {
						writeError(w, commonerrors.ErrUnableToProcess, http.StatusInternalServerError)
						return
					}
					session.Values["user-id"] = user.Actor.ID
					session.Save(r, w)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					w.WriteHeader(http.StatusOK)
				}
			}
		} else {
			writeError(w, commonerrors.ErrFormInvalid, http.StatusBadRequest)
		}
	default:
		writeError(w, errors.New("method not allowed"), http.StatusMethodNotAllowed)
	}
}

// Logout deletes session from redis.
func Logout(w http.ResponseWriter, r *http.Request) {
	session, err := db.Context.Store.Get(r, "web")
	if err != nil {
		writeError(w, commonerrors.ErrUnableToProcess, http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		writeError(w, commonerrors.ErrUnableToProcess, http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func getUserByUsername(username string, domain string) (*models.User, error) {
	if !utils.ContainsString(db.Context.Settings.LocalDomains, domain) {
		return nil, commonerrors.ErrIsRemote
	}
	var user models.User
	err := db.DB().Joins("Actor").Where(`lower("Actor"."username") = lower(?)`, username).Where(`lower("Actor"."domain") = lower(?)`, domain).First(&user).Error
	return &user, err
}

func getUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB().Where("lower(email) = lower(?)", email).Joins("Actor").First(&user).Error
	return &user, err
}
