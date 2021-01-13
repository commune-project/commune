package models

import "github.com/commune-project/commune/models/abstract"

// Commune is the actor to relay the posts.
type Commune struct {
	abstract.Actor
	Categories []Category
	Members    []*Account `gorm:"many2many:commune_members;"`
}
