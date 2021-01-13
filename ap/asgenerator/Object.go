package asgenerator

import (
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/utils"
)

// GenerateAS generates the AS JSON map from an object in models.
func GenerateAS(obj interface{}) map[string]interface{} {
	if value, ok := obj.(interfaces.IActor); ok {
		return generateASIActor(value)
	} else if value, ok := obj.(interfaces.IPost); ok {
		return generateASIPost(value)
	} else if value, ok := obj.(interfaces.IObject); ok {
		return generateASIObject(value)
	} else {
		return map[string]interface{}{}
	}
}

func generateASIObject(obj interfaces.IObject) map[string]interface{} {
	return utils.ConcatMaps([]map[string]interface{}{
		{
			"id":        obj.GetURI(),
			"url":       obj.GetURL(),
			"type":      obj.GetType(),
			"published": obj.GetPublished(),
			"updated":   obj.GetUpdated(),
		},
		obj.RestChildren(),
	})
}
