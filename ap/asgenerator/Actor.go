package asgenerator

import (
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/utils"
)

func generateASIActor(obj interfaces.IActor) map[string]interface{} {
	return utils.ConcatMaps([]map[string]interface{}{
		generateASIObject(obj),
		{
			"preferredUsername": obj.GetUsername(),
			"inbox":             obj.GetInboxURI(),
			"outbox":            obj.GetOutboxURI(),
			"following":         obj.GetFollowingURI(),
			"followers":         obj.GetFollowersURI(),
			"publicKey": map[string]interface{}{
				"id":           obj.GetPublicKeyURI(),
				"owner":        obj.GetURI(),
				"publicKeyPem": obj.GetPublicKey(),
			},
		},
		obj.RestChildren(),
	})
}
