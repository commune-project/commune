package interfaces

type IActor interface {
	IObject
	GetUsername() string
	GetPublicKey() string
	GetPublicKeyURI() string
	GetFollowersURI() string
	GetFollowingURI() string
	GetInboxURI() string
	GetOutboxURI() string
	IsBot() bool
	IsCommune() bool
}

var UserTypes []string = []string{"Person", "Service"}
var ActorTypes []string = []string{"Person", "Service", "Application", "Group"}
