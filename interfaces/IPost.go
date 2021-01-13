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
