package validators

import (
	"github.com/commune-project/commune/interfaces"
)

// ValidatePost checks a post's properties
func ValidatePost(post interfaces.IPost) error {
	validators := [](func(post interfaces.IPost) error){}
	for _, validator := range validators {
		if err := validator(post); err != nil {
			return err
		}
	}
	return nil
}

func validatePostURI(post interfaces.IPost) error {
	if uri := post.GetURI(); uri == "" {
		return &validatingError{
			kind:  post.GetType(),
			key:   "id",
			value: uri,
		}
	}
	return nil
}

func validatePostURLs(post interfaces.IPost) error {
	m := map[string](func() string){
		"id":  post.GetURI,
		"url": post.GetURL,
	}
	for k, v := range m {
		value := v()
		if value != "" && !isURL(value) {
			return &validatingError{
				kind:  post.GetType(),
				key:   k,
				value: value,
			}
		}
	}
	return nil
}
