package asgenerator

import (
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/utils"
)

func generateASIPost(obj interfaces.IPost) map[string]interface{} {
	return utils.ConcatMaps([]map[string]interface{}{
		generateASIObject(obj),
		{
			"attributedTo": obj.GetAuthorURI(),
			"name":         obj.GetName(),
			"content":      obj.GetContent(),
			"mediaType":    "text/html",
			"source": map[string]interface{}{
				"content":   obj.GetSource(),
				"mediaType": obj.GetSourceMediaType(),
			},
			"inReplyTo": obj.GetInReplyTo(),
		},
		obj.RestChildren(),
	})
}
