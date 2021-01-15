package db

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/models/abstract"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
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
	{
		var b [8]byte
		_, err := crypto_rand.Read(b[:])
		if err != nil {
			rand.Seed(time.Now().Unix())
		} else {
			rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
		}
	}
	if err := DB.Exec(fmt.Sprintf(`CREATE OR REPLACE FUNCTION public.timestamp_id_secure_random_hex() RETURNS text LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE
	AS $BODY$
	BEGIN
		RETURN SUBSTRING(md5(%d::text), 1, 16);
  	END
  	$BODY$`, rand.Int63())).Error; err != nil {
		fmt.Println(err)
	}
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
	DB.Create(&models.Post{
		Object: abstract.Object{
			Type: "Note",
		},
		Author:          user.Account,
		Content:         "什么鬼",
		MediaType:       "text/plain",
		Source:          "什么鬼",
		SourceMediaType: "text/plain",
		Name:            "说说",
		ReplyTo:         nil,
	})
}
