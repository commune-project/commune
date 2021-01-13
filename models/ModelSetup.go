package models

import "gorm.io/gorm"

// ModelSetup creates join tables etc.
func ModelSetup(db *gorm.DB) {
	db.SetupJoinTable(&Account{}, "JoinedCommunes", &CommuneMember{})
	db.SetupJoinTable(&Commune{}, "Members", &CommuneMember{})
}
