package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/handlers"
	"github.com/commune-project/commune/handlers/middleware"
	"github.com/commune-project/commune/webfinger"
	"github.com/gorilla/mux"
)

var globalRouter *mux.Router

func init() {
	globalRouter = mux.NewRouter()
	setupRouter(globalRouter)
}

// GetRouter returns a configurated mux.Router
func GetRouter() *mux.Router {
	return globalRouter
}

func setupRouter(router *mux.Router) {
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/api/commune/login", handlers.Login)
	router.HandleFunc("/api/commune/logout", handlers.Logout)
	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/.well-known/webfinger", webfinger.Handler)
	apSubRouter := router.PathPrefix("/").Subrouter()

	// Users
	apSubRouter.HandleFunc("/users/{username}", handlers.ApUserHandler)
	apSubRouter.HandleFunc("/users/{username}/outbox", handlers.ApUserOutboxHandler)

	// Communities
	apSubRouter.HandleFunc("/communities/{username}", handlers.ApUserHandler)
	apSubRouter.HandleFunc("/communities/{username}/outbox", handlers.ApUserOutboxHandler)

	// Posts
	apSubRouter.HandleFunc("/p/{id:[0-9]+}", handlers.ApPostHandler)
	apSubRouter.HandleFunc("/p/{id:[0-9]+}/activity", handlers.ApPostActivityHandler)
	apSubRouter.HandleFunc("/users/{username}/statuses/{id:[0-9]+}", handlers.ApPostHandler)
	apSubRouter.HandleFunc("/users/{username}/statuses/{id:[0-9]+}/activity", handlers.ApPostActivityHandler)

	apSubRouter.Use(mux.MiddlewareFunc(middleware.Authenticate(log.New(log.Writer(), "auth: ", 0), db.Context)))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category: qunimade\n")
}
