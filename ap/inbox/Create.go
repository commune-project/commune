package inbox

import (
	"net/http"

	"github.com/commune-project/commune/ap/asparser"
	"github.com/commune-project/commune/ap/fetchers"
	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/interfaces/validators"
	"github.com/commune-project/commune/logic/posts"
	"github.com/commune-project/commune/models"
	"github.com/commune-project/commune/utils/commonerrors"
)

type createHandler struct{}

func (handler createHandler) Process(r *http.Request, data map[string]interface{}, processingInfo *ProcessingInfo) (err error) {
	var ipost interfaces.IPost

	if object, ok := data["object"].(map[string]interface{}); ok {
		// Check if already exists.
		{
			objectURI, ok := object["id"].(string)
			if !ok {
				return commonerrors.ErrFormInvalid
			}
			post, err := dbmanagers.GetPostByURI(db.Context, objectURI)
			if post != nil && err == nil {
				return commonerrors.ErrAlreadyExists
			}
		}
		ipost = asparser.ParseIPost(object)
	} else if objectURI, ok := data["object"].(string); ok {
		// Check if already exists.
		{
			post, err := dbmanagers.GetPostByURI(db.Context, objectURI)
			if post != nil && err == nil {
				return commonerrors.ErrAlreadyExists
			}
		}
		ipost, err = fetchers.FetchPostByURI(db.Context, objectURI)
		if err != nil {
			return err
		}
	}

	err = validators.ValidatePost(ipost)
	if err != nil {
		return
	}

	if post, ok := ipost.(*models.Post); ok {
		err = posts.CreatePost(db.Context, post)
		return
	}

	return commonerrors.ErrUnableToProcess
}
