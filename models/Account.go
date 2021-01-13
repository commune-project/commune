package models

import "github.com/commune-project/commune/models/abstract"

// Account is a user controlled by either a person or a bot.
type Account struct {
	abstract.Actor
	Posts          []Post     `gorm:"foreignKey:AuthorID"`
	JoinedCommunes []*Commune `gorm:"many2many:commune_members;"`
}
