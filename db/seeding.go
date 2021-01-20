package db

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"

	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/models/abstract"
	"gorm.io/gorm"
)

// DefaultPassword sets the password for testing. DO NOT USE IT IN PRODUCTION!
const DefaultPassword string = "123456"

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
	if err := DB().Exec(fmt.Sprintf(`CREATE OR REPLACE FUNCTION public.timestamp_id_secure_random_hex() RETURNS text LANGUAGE plpgsql IMMUTABLE PARALLEL SAFE
	AS $BODY$
	BEGIN
		RETURN SUBSTRING(md5(%d::text), 1, 16);
  	END
  	$BODY$`, rand.Int63())).Error; err != nil {
		panic(err)
	}
	seed("commune1.localdomain")
	seed("commune1.misaka4e21.science")

}

func seed(domain string) {
	user := models.NewUser("misaka4e22", domain, "misaka4e21@"+domain, DefaultPassword)
	if user == nil {
		panic("unable to create user.")
	}
	err := DB().Create(user).Error
	if err != nil {
		panic(err)
	}

	var defaultCategory *models.Category
	err = DB().Transaction(func(tx *gorm.DB) error {
		localCommune := models.NewLocalCommune("limelight", domain, true)

		if err := tx.Create(localCommune).Error; err != nil {
			return err
		}

		communeMembership := &models.Follow{
			Following: localCommune.Actor,
			Follower:  user.Actor,
			Role:      "creator",
		}
		if err := tx.Create(communeMembership).Error; err != nil {
			return err
		}
		defaultCategory = &models.Category{
			Commune: &localCommune.Actor,
			Slug:    "default",
		}
		return tx.Create(defaultCategory).Error
	})
	if err != nil {
		panic(err)
	}

	DB().Create(&models.Post{
		Object: abstract.Object{
			Type: "Note",
		},
		Author:          user.Actor,
		Category:        defaultCategory,
		Content:         "什么鬼",
		MediaType:       "text/plain",
		Source:          "什么鬼",
		SourceMediaType: "text/plain",
		Name:            "去你妈的鬼神",
		ReplyTo:         nil,
	})
	DB().Create(&models.Post{
		Object: abstract.Object{
			Type: "Note",
		},
		Author:          user.Actor,
		Content:         "什么鬼",
		MediaType:       "text/plain",
		Source:          "什么鬼",
		SourceMediaType: "text/plain",
		Name:            "说说",
		ReplyTo:         nil,
	})
}
