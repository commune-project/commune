package models_test

import (
	"testing"

	"github.com/commune-project/commune/models"
)

func Test_NewUser(t *testing.T) {
	if user, err := models.NewUser("misaka4e21", "m.hitorino.moe", "misaka4e21@rbq.press", "123456"); err != nil {
		t.Error(err)
	} else {
		if uri := user.Account.GetURI(); uri != "https://m.hitorino.moe/users/misaka4e21" {
			t.Error("user.Account.GetURI() wrong: " + uri)
		}
		if url := user.Account.GetURL(); url != "https://m.hitorino.moe/@misaka4e21" {
			t.Error("user.Account.GetURL() wrong: " + url)
		}
		if uri := user.Account.GetInboxURI(); uri != "https://m.hitorino.moe/users/misaka4e21/inbox" {
			t.Error("user.Account.GetInboxURI() wrong: " + uri)
		}
		if uri := user.Account.GetOutboxURI(); uri != "https://m.hitorino.moe/users/misaka4e21/outbox" {
			t.Error("user.Account.GetOutboxURI() wrong: " + uri)
		}
		if uri := user.Account.GetFollowersURI(); uri != "https://m.hitorino.moe/users/misaka4e21/followers" {
			t.Error("user.Account.GetFollowersURI() wrong: " + uri)
		}
		if uri := user.Account.GetFollowingURI(); uri != "https://m.hitorino.moe/users/misaka4e21/following" {
			t.Error("user.Account.GetFollowingURI() wrong: " + uri)
		}
		t.Log("All URIs were tested.")
		if !user.Account.IsLocal([]string{"m.hitorino.moe"}) {
			t.Error("Local user determined as remote.")
		}
	}
}
