package router

import (
	"fmt"
	"net/http"

	"github.com/commune-project/commune/handlers"
	"github.com/gorilla/mux"
)

// GetRouter returns a configurated mux.Router
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	setupRouter(r)
	return r
}

func setupRouter(router *mux.Router) {
	router.HandleFunc("/", homeHandler)

	// Users
	router.HandleFunc("/users/{username}", handlers.ApUserHandler)
	router.HandleFunc("/users/{username}/outbox", handlers.ApUserOutboxHandler)

	// Communities
	router.HandleFunc("/communities/{username}", handlers.ApCommunityHandler)
	router.HandleFunc("/communities/{username}/outbox", handlers.ApCommunityOutboxHandler)

	// Posts
	router.HandleFunc("/p/{id:[0-9]+}", handlers.ApPostHandler)
	router.HandleFunc("/p/{id:[0-9]+}/activity", handlers.ApPostActivityHandler)
	router.HandleFunc("/users/{username}/statuses/{id:[0-9]+}", handlers.ApPostHandler)
	router.HandleFunc("/users/{username}/statuses/{id:[0-9]+}/activity", handlers.ApPostActivityHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: qunimade\n")
}
