package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/models/abstract"
	"github.com/commune-project/commune/router"
	"gorm.io/gorm"
)

func main() {
	fmt.Printf("Saluton mondo.\n")
	db.DB.AutoMigrate(&models.Account{}, &models.User{}, &models.CommuneMember{}, &models.Category{}, &models.Commune{}, &models.Post{})

	fmt.Println("???")
	db.DB.Exec("DELETE FROM posts")
	db.DB.Exec("DELETE FROM users")
	db.DB.Exec("DELETE FROM accounts")
	user, err := models.NewUser("misaka4e22", "m.hitorino.moe", "misaka4e21@gmail.com", "123456")
	if err != nil {
		panic(err)
	}
	db.DB.Create(user)

	db.DB.Transaction(func(tx *gorm.DB) error {
		localCommune, _ := models.NewLocalCommune("limelight", "m.hitorino.moe")
		if err := tx.Create(localCommune).Error; err != nil {
			return err
		}

		communeMembership := &models.CommuneMember{
			Commune: localCommune.Commune,
			Account: user.Account,
			Role:    "creator",
		}
		return db.DB.Create(communeMembership).Error

	})

	db.DB.Create(&models.Post{
		Object: abstract.Object{
			Type: "Note",
		},
		Author:          user.Account,
		Category:        nil,
		Content:         "什么鬼",
		MediaType:       "text/plain",
		Source:          "什么鬼",
		SourceMediaType: "text/plain",
		Name:            "去你妈的鬼神",
		ReplyTo:         nil,
	})

	srv := &http.Server{
		Handler: router.GetRouter(),
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
