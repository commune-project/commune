package inbox

import (
	"net/http"

	"github.com/commune-project/commune/models"
)

// IInboxHandler describes a generic way to convert a JSON map into a database object.
type IInboxHandler interface {
	Process(r *http.Request, data map[string]interface{}, processingInfo *ProcessingInfo) error
}

// ProcessingInfo stores the intermediate information for processing.
type ProcessingInfo struct {
	Actor           *models.Actor
	Post            *models.Post
	AdditionalPosts []models.Post
}

type sequenceHandler struct {
	handlers []IInboxHandler
}

func (thisHandler sequenceHandler) Process(r *http.Request, data map[string]interface{}, processingInfo *ProcessingInfo) error {
	for _, handler := range thisHandler.handlers {
		err := handler.Process(r, data, processingInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func mapGetString(data map[string]interface{}, key string) (string, bool) {
	if value, ok := data[key]; ok {
		if str, ok := value.(string); ok {
			if str != "" {
				return str, true
			}
		}
	}
	return "", false
}
