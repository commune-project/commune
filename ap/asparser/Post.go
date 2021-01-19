package asparser

import (
	"time"

	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/models/abstract"
	"github.com/commune-project/commune/utils"
)

// // Post is a general posting.
// type Post struct {
// 	abstract.Object
// 	AuthorID        int
// 	Author          Actor
// 	CategoryID      *int
// 	Category        *Category
// 	Content         string
// 	MediaType       string
// 	Source          string
// 	SourceMediaType string
// 	Name            string
// 	ReplyToID       *int
// 	ReplyTo         *Post
// 	BoostToID       *int
// 	BoostTo         *Post
// }

func ParseIPost(data map[string]interface{}) interfaces.IPost {
	if dataType, ok := data["type"].(string); ok && utils.ContainsString(interfaces.PostTypes, dataType) {
		return parseIntoPost(data)
	}
	return nil
}

func parseIntoPost(data map[string]interface{}) *models.Post {
	dataURI, _ := data["id"].(string)
	dataURL, _ := data["url"].(string)
	dataType, _ := data["type"].(string)
	dataName, _ := data["name"].(string)
	dataContent, _ := data["content"].(string)
	dataSource, _ := data["source"].(map[string]interface{})
	dataSourceContent, _ := dataSource["content"].(string)
	dataSourceMediaType, _ := dataSource["mediaType"].(string)
	dataPublished, _ := dataSource["published"].(string)
	createdAt, _ := time.Parse(time.RFC3339, dataPublished)
	// dataSummary, _ := data["summary"].(string)

	return &models.Post{
		Object: abstract.Object{
			Model: abstract.Model{
				CreatedAt: createdAt,
			},
			Type: dataType,
			URI:  &dataURI,
			URL:  dataURL,
		},
		Name:    dataName,
		Content: dataContent,
		// Summary:      dataSummary,
		Source:          dataSourceContent,
		SourceMediaType: dataSourceMediaType,
	}
}
