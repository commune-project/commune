package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/router"
)

func main() {
	fmt.Printf("Saluton mondo.\n")
	models.DropTables(db.DB)
	models.Migrate(db.DB)
	db.Seeding()

	srv := &http.Server{
		Handler: router.GetRouter(),
		Addr:    "0.0.0.0:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println(srv.Addr)

	log.Fatal(srv.ListenAndServe())
}
