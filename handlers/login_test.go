package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/commune-project/commune/router"
	"github.com/commune-project/commune/utils/commonerrors"
)

func TestLogin(t *testing.T) {
	actor, err := dbmanagers.GetActorByUsername(db.DB(), "misaka4e22", "commune1.localdomain")
	if err != nil {
		t.Error("Unable to get Actor: ", err, " Check your database.")
	}

	_, w := doLogin("misaka4e22", db.DefaultPassword)(nil)
	if w.Code != 200 {
		t.Error(commonerrors.ErrNotLoggedIn, ": ", string(w.Body.Bytes()))
	}
	// Test fetching homepage with returned cookies
	r2, _ := http.NewRequest("GET", "https://commune1.localdomain/", nil)
	for _, cookie := range w.Result().Cookies() {
		r2.AddCookie(cookie)
	}
	session, err := db.Context.Store.Get(r2, "web")
	if err != nil {
		t.Error("Unable to get session store: ", err, " Check your redis.")
	}
	if session.Values["user-id"] != actor.ID {
		t.Error("Session id ", session.Values["user-id"], "should be: ", actor.ID)
	}
}

func TestLoginWithEmail(t *testing.T) {
	actor, err := dbmanagers.GetActorByUsername(db.DB(), "misaka4e22", "commune1.localdomain")
	if err != nil {
		t.Error("Unable to get Actor: ", err, " Check your database.")
	}

	_, w := doLogin("misaka4e21@commune1.localdomain", db.DefaultPassword)(nil)
	if w.Code != 200 {
		t.Error(commonerrors.ErrNotLoggedIn, ": ", string(w.Body.Bytes()))
	}
	// Test fetching homepage with returned cookies
	r2, _ := http.NewRequest("GET", "https://commune1.localdomain/", nil)
	for _, cookie := range w.Result().Cookies() {
		r2.AddCookie(cookie)
	}
	session, err := db.Context.Store.Get(r2, "web")
	if err != nil {
		t.Error("Unable to get session store: ", err, " Check your redis.")
	}
	if session.Values["user-id"] != actor.ID {
		t.Error("Session id ", session.Values["user-id"], "should be: ", actor.ID)
	}
}

func TestLoginWithWrongUsername(t *testing.T) {
	_, w := doLogin("misaka4e23", "???")(nil)
	if w.Code == 200 {
		t.Error("There is no user misaka4e23, shouldn't be allowed to login.")
	}
}

func TestLogout(t *testing.T) {
	_, w := doLogin("misaka4e22", db.DefaultPassword)(nil)
	_, w = doLogout(w.Result().Cookies())

	r, _ := http.NewRequest("GET", "https://commune1.localdomain/", nil)

	session, _ := db.Context.Store.Get(r, "web")
	if actorID, ok := session.Values["user-id"].(int); ok {
		t.Error("Session shouldn't exist:", actorID)
	}
}

func doLogin(username string, password string) func(cookies []*http.Cookie) (*http.Request, *httptest.ResponseRecorder) {
	// doLogin
	return func(cookies []*http.Cookie) (*http.Request, *httptest.ResponseRecorder) {
		// Request y writer for testing
		bodyReader := strings.NewReader(fmt.Sprintf(`{
			"username": "%s",
			"password": "%s"
		}`, username, password))
		r, _ := http.NewRequest("POST", "https://commune1.localdomain/api/commune/login", bodyReader)
		r.Header.Add("Content-Type", "application/json")
		for _, cookie := range cookies {
			r.AddCookie(cookie)
		}
		w := httptest.NewRecorder()

		router.GetRouter().ServeHTTP(w, r)

		return r, w
	}
}

func doLogout(cookies []*http.Cookie) (*http.Request, *httptest.ResponseRecorder) {
	// Request y writer for testing
	r, _ := http.NewRequest("POST", "https://commune1.localdomain/api/commune/logout", nil)
	w := httptest.NewRecorder()
	for _, cookie := range cookies {
		r.AddCookie(cookie)
	}

	router.GetRouter().ServeHTTP(w, r)

	return r, w
}
