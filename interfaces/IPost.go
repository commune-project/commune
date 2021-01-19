package interfaces

type IPost interface {
	IObject
	GetName() string
	GetContent() string
	GetSource() string
	GetSourceMediaType() string
	GetAuthorURI() string
	GetActivityCreate() IObject
	GetInReplyTo() string
}

var PostTypes []string = []string{"Note", "Article", "Page"}
