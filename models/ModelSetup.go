package models

import "gorm.io/gorm"

// ModelSetup creates join tables etc.
func ModelSetup(db *gorm.DB) {
	db.SetupJoinTable(&Account{}, "JoinedCommunes", &CommuneMember{})
	db.SetupJoinTable(&Commune{}, "Members", &CommuneMember{})
}

// Migrate does gorm automatic migrations.
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Account{}, &User{}, &CommuneMember{}, &LocalCommune{}, &Category{}, &Commune{}, &Post{})
}

// DeleteFromTables removes all contents in the database!
func DeleteFromTables(db *gorm.DB) {
	db.Delete(&Account{}, "1=1")
	db.Delete(&User{}, "1=1")
	db.Delete(&CommuneMember{}, "1=1")
	db.Delete(&LocalCommune{}, "1=1")
	db.Delete(&Category{}, "1=1")
	db.Delete(&Commune{}, "1=1")
	db.Delete(&Post{}, "1=1")
}
