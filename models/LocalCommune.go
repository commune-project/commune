package models

import (
	"github.com/commune-project/commune/models/abstract"
	"github.com/commune-project/commune/utils"
)

// LocalCommune is the private part of a local commune.
type LocalCommune struct {
	CommuneID  int     `gorm:"primaryKey"`
	Commune    Commune `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	PrivateKey string
	IsVerified bool
}

// NewLocalCommune creates a new User struct without insert it into database.
func NewLocalCommune(username string, domain string) (commune *LocalCommune, err error) {
	publicKeyPEM, privateKeyPEM := utils.GenerateRsaKeys()
	return &LocalCommune{
		Commune: Commune{
			Actor: abstract.Actor{
				Object: abstract.Object{
					Type: "Group",
				},
				Username:  username,
				Domain:    domain,
				PublicKey: string(publicKeyPEM),
			},
		},
		PrivateKey: string(privateKeyPEM),
		IsVerified: true,
	}, nil
}
