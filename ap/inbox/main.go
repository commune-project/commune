package inbox

import (
	"net/http"
)

// IInboxHandler describes a generic way to convert a JSON map into a database object.
type IInboxHandler interface {
	Process(r *http.Request, data map[string]interface{}, object interface{}) error
}

type inboxHandler struct {
	handlers []IInboxHandler
}

func (thisHandler inboxHandler) Process(r *http.Request, data map[string]interface{}, object interface{}) error {
	for _, handler := range thisHandler.handlers {
		err := handler.Process(r, data, object)
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
