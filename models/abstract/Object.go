package abstract

import (
	"net/url"
	"time"

	"github.com/commune-project/commune/interfaces"
	"gorm.io/gorm"
)

// Model is a gorm.Model with signed PK.
type Model struct {
	ID        int            `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time      `json:"published" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Object is just a template, don't create an actual table for it.
type Object struct {
	Model
	URI  *string `gorm:"unique" json:"id"`
	URL  *string `json:"url"`
	Type string  `json:"type"`
}

func (obj *Object) GetDomain() string {
	u, err := url.Parse(*obj.URI)
	if err != nil {
		return ""
	} else {
		return u.Host
	}
}

func (obj *Object) GetURI() string {
	if obj.URI == nil {
		return ""
	}
	return *obj.URI
}

func (obj *Object) GetURL() string {
	if obj.URL == nil {
		return ""
	}
	return *obj.URL
}

func (obj *Object) IsLocal(localDomains []string) bool {
	var iobj interfaces.IObject = obj
	objDomain := iobj.GetDomain()
	for _, localDomain := range localDomains {
		if objDomain == localDomain {
			return true
		}
	}
	return false
}

func (obj *Object) GetType() string {
	return obj.Type
}

func (obj *Object) GetPublished() string {
	return obj.CreatedAt.UTC().Format(time.RFC3339)
}

func (obj *Object) GetUpdated() string {
	return obj.UpdatedAt.UTC().Format(time.RFC3339)
}

func (obj *Object) RestChildren() map[string]interface{} {
	return map[string]interface{}{}
}
