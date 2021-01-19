package inbox

import (
	"net/http"

	"github.com/commune-project/commune/ap/asparser"
	"github.com/commune-project/commune/ap/fetchers"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/interfaces/validators"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/utils/commonerrors"
)

type createHandler struct{}

func (handler createHandler) Process(r *http.Request, data map[string]interface{}, processingInfo *ProcessingInfo) (err error) {
	var ipost interfaces.IPost

	if object, ok := data["object"].(map[string]interface{}); ok {
		ipost = asparser.ParseIPost(object)
	} else if objectURI, ok := data["object"].(string); ok {
		ipost, err = fetchers.GetOrFetchPostByURI(db.Context, objectURI)
		if err != nil {
			return err
		}
	}

	err = validators.ValidatePost(ipost)
	if err != nil {
		return
	}

	if post, ok := ipost.(*models.Post); ok {
		err = db.DB().Create(post).Error
		return
	}

	return commonerrors.ErrUnableToProcess
}
