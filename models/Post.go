package models

import (
	"fmt"

	"github.com/commune-project/commune/ap/asgenerator"
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/models/abstract"
)

// Post is a general posting.
type Post struct {
	abstract.Object

	AuthorID        int
	Author          Account
	CategoryID      *int
	Category        *Category
	Content         string
	MediaType       string
	Source          string
	SourceMediaType string
	Name            string
	ReplyToID       *int
	ReplyTo         *Post
}

func (post *Post) GetDomain() string {
	if post.Category == nil {
		return post.Author.GetDomain()
	}
	return post.Category.GetDomain()
}

func (post *Post) GetURI() string {
	if post.URI == nil {
		return fmt.Sprintf("https://%s/p/%d", post.GetDomain(), post.ID)
	}
	return *post.URI
}

func (post *Post) GetURL() string {
	if post.URL == nil {
		return post.GetURI()
	}
	return *post.URL
}

func (post *Post) GetName() string {
	return post.Name
}

func (post *Post) GetContent() string {
	return post.Content
}

func (post *Post) GetSource() string {
	return post.Source
}

func (post *Post) GetSourceMediaType() string {
	return post.SourceMediaType
}

func (post *Post) GetAuthorURI() string {
	return post.Author.GetURI()
}

func (post *Post) GetActivityCreate() interfaces.IObject {
	return postActivityCreate{
		post: post,
	}
}

func (post *Post) GetInReplyTo() string {
	if parent := post.ReplyTo; parent != nil {
		return parent.GetURI()
	}
	return ""
}

type postActivityCreate struct {
	post *Post
}

func (this postActivityCreate) GetDomain() string {
	return this.post.GetDomain()
}
func (this postActivityCreate) GetURI() string {
	return this.post.GetURI() + "/activity"
}

func (this postActivityCreate) GetURL() string {
	return this.post.GetURL()
}

func (this postActivityCreate) GetType() string {
	return "Create"
}

func (this postActivityCreate) IsLocal(localDomains []string) bool {
	return this.post.IsLocal(localDomains)
}

func (this postActivityCreate) RestChildren() map[string]interface{} {
	return map[string]interface{}{
		"actor":  this.post.Author.GetURI(),
		"object": asgenerator.GenerateAS(this.post),
	}
}

func (this postActivityCreate) GetPublished() string {
	return this.post.GetPublished()
}

func (this postActivityCreate) GetUpdated() string {
	return this.post.GetPublished()
}
