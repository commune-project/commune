package asgenerator_test

import (
	"testing"

	"github.com/commune-project/commune/ap/asgenerator"
	"github.com/commune-project/commune/interfaces"
	"github.com/commune-project/commune/models"
)

func Test_Generate_Actor(t *testing.T) {
	user := models.NewUser("misaka4e22", "m.hitorino.moe", "misaka4e21@gmail.com", "123456")
	if user == nil {
		t.Error("Unable to create user.")
	}
	var account interfaces.IActor = &user.Actor
	mjson := asgenerator.GenerateAS(account)

	if mjson["preferredUsername"] != "misaka4e22" {
		t.Error("preferredUsername wrong")
	}
	if publicKey, ok := mjson["publicKey"].(map[string]interface{}); ok {
		if publicKey["id"] != "https://m.hitorino.moe/users/misaka4e22#main-key" {
			t.Error("publicKey id wrong.")
		}
	} else {
		t.Error("publicKey isn't a map.")
	}
}
