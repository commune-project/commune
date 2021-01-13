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
