package router

import (
	"fmt"
	"net/http"

	"github.com/commune-project/commune/handlers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	SetupRouter(r)
	return r
}

func SetupRouter(router *mux.Router) {
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/users/{username}", handlers.ApUserHandler)
	router.HandleFunc("/users/{username}/outbox", handlers.ApUserOutboxHandler)
	router.HandleFunc("/p/{id:[0-9]+}", handlers.ApPostHandler)
	router.HandleFunc("/p/{id:[0-9]+}/activity", handlers.ApPostActivityHandler)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: qunimade\n")
}
