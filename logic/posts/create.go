package posts

import (
	"regexp"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/microcosm-cc/bluemonday"
)

func CreatePost(context db.SiteContext, post *models.Post) error {
	convertSource(context, post)
	sanitize(post)
	return context.DB.Create(post).Error
}

func convertSource(context db.SiteContext, post *models.Post) {
	if !post.IsLocal(context.Settings.LocalDomains) {
		return
	}
	post.Content = ""
	switch post.SourceMediaType {
	case "text/html":
		post.Content = post.Source
	}
}

func sanitize(post *models.Post) {
	p := bluemonday.NewPolicy()
	p.AllowElements("p", "span", "br", "a")
	p.AllowElements("h1", "h2", "h3", "h4", "h5", "h6", "blockquote", "pre", "ul", "ol", "li", "table", "tr", "th", "td")
	p.AllowAttrs("class").Matching(regexp.MustCompile(`^(h-[^\s]+|p-[^\s]+|u-[^\s]+|dt-[^\s]+|e-[^\s]+|mention|hashtag|ellipsis|invisible)\s?+$`))
	p.AllowImages()
	p.AllowStandardAttributes()

	post.Content = p.Sanitize(post.Content)
}
