package models

import (
	"fmt"

	"github.com/commune-project/commune/models/abstract"
)

// Category contains many Posts.
type Category struct {
	abstract.Object
	CommuneID int
	Commune   *Commune
	Slug      string
	Posts     []*Post
}

func (category *Category) GetDomain() string {
	if category == nil {
		return ""
	} else if category.Commune != nil {
		return category.Commune.GetDomain()
	}
	return category.Object.GetDomain()
}

func (category *Category) GetURI() string {
	if category.URI != nil {
		return *category.URI
	}
	return fmt.Sprintf("%s/categories/%s", category.Commune.GetURI(), category.Slug)
}

func (category *Category) GetURL() string {
	if category.URL != nil {
		return *category.URL
	}
	return fmt.Sprintf("%s/c/%s", category.Commune.GetURL(), category.Slug)
}
