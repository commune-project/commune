package models

import (
	"github.com/commune-project/commune/models/abstract"
	"github.com/commune-project/commune/utils"
	"golang.org/x/crypto/bcrypt"
)

// User is the private part of a local account.
type User struct {
	ActorID           int     `gorm:"primaryKey;autoIncrement:false;"`
	Actor             Actor   `gorm:"foreignKey:ActorID"`
	Email             *string `gorm:"unique"`
	EncryptedPassword string
	PrivateKey        string
	IsVerified        bool
	IsAdmin           bool
	IsModerator       bool
}

// NewUser creates a new User struct without insert it into database.
func NewUser(username string, domain string, email string, password string) *User {
	publicKeyPEM, privateKeyPEM := utils.GenerateRsaKeys()
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return nil
	}
	return &User{
		Actor: Actor{
			Object: abstract.Object{
				// URI:  utils.PS(fmt.Sprintf("https://%s/users/%s", domain, username)),
				// URL:  utils.PS(fmt.Sprintf("https://%s/@%s", domain, username)),
				Type: "Person",
			},
			Username:  username,
			Domain:    domain,
			PublicKey: string(publicKeyPEM),
			// InboxURI:     fmt.Sprintf("https://%s/users/%s/inbox", domain, username),
			// OutboxURI:    fmt.Sprintf("https://%s/users/%s/outbox", domain, username),
			// FollowersURI: fmt.Sprintf("https://%s/users/%s/followers", domain, username),
			// FollowingURI: fmt.Sprintf("https://%s/users/%s/following", domain, username),
		},
		PrivateKey:        string(privateKeyPEM),
		Email:             &email,
		EncryptedPassword: string(encryptedPassword),
		IsVerified:        false,
		IsAdmin:           false,
		IsModerator:       false,
	}
}

// NewLocalCommune creates a new User struct without insert it into database.
func NewLocalCommune(username string, domain string, autoVerify bool) *User {
	publicKeyPEM, privateKeyPEM := utils.GenerateRsaKeys()
	return &User{
		Actor: Actor{
			Object: abstract.Object{
				Type: "Group",
			},
			Username:  username,
			Domain:    domain,
			PublicKey: string(publicKeyPEM),
		},
		PrivateKey:  string(privateKeyPEM),
		IsVerified:  autoVerify,
		IsAdmin:     false,
		IsModerator: false,
	}
}
