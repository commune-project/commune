package models_test

import (
	"testing"

	"github.com/commune-project/commune/models"
)

func Test_NewUser(t *testing.T) {
	if user := models.NewUser("misaka4e21", "m.hitorino.moe", "misaka4e21@rbq.press", "123456"); user == nil {
		t.Error("Unable to create user.")
	} else {
		if uri := user.Actor.GetURI(); uri != "https://m.hitorino.moe/users/misaka4e21" {
			t.Error("user.Actor.GetURI() wrong: " + uri)
		}
		if url := user.Actor.GetURL(); url != "https://m.hitorino.moe/@misaka4e21" {
			t.Error("user.Actor.GetURL() wrong: " + url)
		}
		if uri := user.Actor.GetInboxURI(); uri != "https://m.hitorino.moe/users/misaka4e21/inbox" {
			t.Error("user.Actor.GetInboxURI() wrong: " + uri)
		}
		if uri := user.Actor.GetOutboxURI(); uri != "https://m.hitorino.moe/users/misaka4e21/outbox" {
			t.Error("user.Actor.GetOutboxURI() wrong: " + uri)
		}
		if uri := user.Actor.GetFollowersURI(); uri != "https://m.hitorino.moe/users/misaka4e21/followers" {
			t.Error("user.Actor.GetFollowersURI() wrong: " + uri)
		}
		if uri := user.Actor.GetFollowingURI(); uri != "https://m.hitorino.moe/users/misaka4e21/following" {
			t.Error("user.Actor.GetFollowingURI() wrong: " + uri)
		}
		if !user.Actor.IsLocal([]string{"m.hitorino.moe"}) {
			t.Error("Local user determined as remote.")
		}
	}
}
