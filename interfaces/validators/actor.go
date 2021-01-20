package validators

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/commune-project/commune/interfaces"
)

type validatingError struct {
	kind  string
	key   string
	value string
}

func (err *validatingError) Error() string {
	return fmt.Sprintf("Error validating %s: %s could not be %s.", err.kind, err.key, err.value)
}

// UsernameRe regulates the username rule.
var UsernameRe *regexp.Regexp

func init() {
	var err error
	UsernameRe, err = regexp.Compile(`[a-z0-9_]+([a-z0-9_\.-]+[a-z0-9_]+)?`)
	if err != nil {
		panic(err)
	}
}

// ValidateActor checks an actor's properties
func ValidateActor(actor interfaces.IActor) error {
	validators := [](func(actor interfaces.IActor) error){
		validateActorUsername,
		validateActorURI,
		validateActorURLs,
	}
	for _, validator := range validators {
		if err := validator(actor); err != nil {
			return err
		}
	}
	return nil
}

func validateActorUsername(actor interfaces.IActor) error {
	username := actor.GetUsername()
	if UsernameRe.MatchString(username) {
		return nil
	}
	return &validatingError{
		kind:  "Actor",
		key:   "preferredUsername",
		value: username,
	}
}

func validateActorURI(actor interfaces.IActor) error {
	if uri := actor.GetURI(); uri == "" {
		return &validatingError{
			kind:  "Actor",
			key:   "id",
			value: uri,
		}
	}
	return nil
}

func validateActorURLs(actor interfaces.IActor) error {
	m := map[string](func() string){
		"id":        actor.GetURI,
		"url":       actor.GetURL,
		"followers": actor.GetFollowersURI,
		"following": actor.GetFollowingURI,
		"inbox":     actor.GetInboxURI,
		"outbox":    actor.GetOutboxURI,
	}
	for k, v := range m {
		value := v()
		if value != "" && !isURL(value) {
			return &validatingError{
				kind:  "Actor",
				key:   k,
				value: value,
			}
		}
	}
	return nil
}

func isURL(sURL string) bool {
	_, err := url.Parse(sURL)
	return err == nil
}
