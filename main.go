package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/commune-project/commune/cli/communectl"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/router"
)

func main() {
	if len(os.Args) > 1 {
		communectl.Communectl()
		return
	}
	fmt.Printf("Saluton mondo.\n")
	models.DeleteFromTables(db.DB())
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
