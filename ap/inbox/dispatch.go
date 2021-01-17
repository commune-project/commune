package inbox

import (
	"net/http"

	"github.com/commune-project/commune/utils/commonerrors"
)

type dispatcher struct {
	handlers map[string]IInboxHandler
}

func (dispatch dispatcher) Process(r *http.Request, data map[string]interface{}, processingInfo *ProcessingInfo) error {
	if kind, ok := mapGetString(data, "type"); ok {
		if handler, ok := dispatch.handlers[kind]; ok {
			return handler.Process(r, data, processingInfo)
		}
		return commonerrors.ErrUnableToProcess
	}
	return commonerrors.ErrFormInvalid
}

// Dispatcher processes data by Activity type.
func Dispatcher(handlers map[string]IInboxHandler) IInboxHandler {
	return dispatcher{
		handlers: handlers,
	}
}
