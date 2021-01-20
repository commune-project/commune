package validators_test

import (
	"testing"

	"github.com/commune-project/commune/interfaces/validators"
	"github.com/commune-project/commune/models"
)

func TestNewUserActorShouldPassValidation(t *testing.T) {
	user := models.NewUser("username", "domain", "email@domain", "password")
	err := validators.ValidateActor(&user.Actor)
	if err != nil {
		t.Error(err)
	}
}

func TestNewUserActorWithInvalidUsername(t *testing.T) {
	username := "user$#@^@#^#^&#&@%*$^*name"
	user := models.NewUser(username, "domain", "email@domain", "password")
	err := validators.ValidateActor(&user.Actor)
	if err == nil {
		t.Error(username + " should not be a valid username.")
	}
}
