package interfaces

type IObject interface {
	GetDomain() string
	GetURI() string
	GetURL() string
	GetType() string
	IsLocal(localDomains []string) bool
	RestChildren() map[string]interface{}
	GetPublished() string
	GetUpdated() string
}
