package models

import "gorm.io/gorm"

// ModelSetup creates join tables etc.
func ModelSetup(db *gorm.DB) {
	db.SetupJoinTable(&Actor{}, "Followers", &Follow{})
	db.SetupJoinTable(&Actor{}, "Following", &Follow{})
}

// Migrate does gorm automatic migrations.
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Actor{}, &User{}, &Follow{}, &Category{}, &Post{})
}

// DeleteFromTables removes all contents in the database!
func DeleteFromTables(db *gorm.DB) {
	db.Delete(&Actor{}, "1=1")
	db.Delete(&User{}, "1=1")
	db.Delete(&Follow{}, "1=1")
	db.Delete(&Category{}, "1=1")
	db.Delete(&Post{}, "1=1")
}
