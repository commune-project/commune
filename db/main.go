package db

import (
	"log"
	"os"
	"strings"

	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/models/abstract"
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
		LocalDomains = []string{}
	} else {
		LocalDomains = strings.Split(sLocalDomains, " ")
	}
}

// Seeding the database. #TODO
func Seeding() {
	user, err := models.NewUser("misaka4e22", "commune1.localdomain", "misaka4e21@gmail.com", "123456")
	if err != nil {
		panic(err)
	}
	err = DB.Create(user).Error
	if err != nil {
		panic(err)
	}

	var defaultCategory *models.Category
	err = DB.Transaction(func(tx *gorm.DB) error {
		localCommune, _ := models.NewLocalCommune("limelight", "commune1.localdomain")
		if err := tx.Create(localCommune).Error; err != nil {
			return err
		}

		communeMembership := &models.CommuneMember{
			Commune: localCommune.Commune,
			Account: user.Account,
			Role:    "creator",
		}
		if err := tx.Create(communeMembership).Error; err != nil {
			return err
		}
		defaultCategory = &models.Category{
			Commune: &localCommune.Commune,
			Slug:    "default",
		}
		return tx.Create(defaultCategory).Error
	})
	if err != nil {
		panic(err)
	}

	DB.Create(&models.Post{
		Object: abstract.Object{
			Type: "Note",
		},
		Author:          user.Account,
		Category:        defaultCategory,
		Content:         "什么鬼",
		MediaType:       "text/plain",
		Source:          "什么鬼",
		SourceMediaType: "text/plain",
		Name:            "去你妈的鬼神",
		ReplyTo:         nil,
	})
}
