package db

import (
	"log"
	"os"
	"strings"

	"github.com/commune-project/commune/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a gorm.DB instance.
var DB *gorm.DB

// LocalDomains contains all domains this program serves.
var LocalDomains []string

func init() {
	openDB()
	readSettings()
}

func openDB() {
	dbURL, isPresent := os.LookupEnv("DATABASE_URL")
	if !isPresent {
		log.Fatal("Please set DATABASE_URL environment var!")
	}
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	models.ModelSetup(db)
	DB = db
}

func readSettings() {
	sLocalDomains, isPresent := os.LookupEnv("COMMUNE_LOCAL_DOMAINS")
	if !isPresent {
		LocalDomains = []string{"commune.example.org"}
	} else {
		LocalDomains = strings.Split(sLocalDomains, " ")
	}
}
