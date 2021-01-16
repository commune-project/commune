package dbmanagers_test

import (
	"testing"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
	"github.com/commune-project/commune/models"
)

func Test_GetActorByUsername(t *testing.T) {
	Actor, err := dbmanagers.GetActorByUsername(db.DB(), "misaka4e22", "commune1.localdomain")
	if err != nil {
		t.Error(err)
	}
	if Actor == nil {
		t.Error("Actor is nil")
	}
}

func Test_GetActorByUsernameWithDifferentDomain(t *testing.T) {
	Actor, err := dbmanagers.GetActorByUsername(db.DB(), "misaka4e22", "not.exist.localdomain")
	if err == nil && Actor != nil {
		t.Error(Actor.GetUsername() + "@" + Actor.GetDomain() + "shouldn't exist")
	}
}

func Test_GetActorByURIWithLocalDomain(t *testing.T) {
	actor, err := dbmanagers.GetActorByURI(db.Context, "https://commune1.localdomain/users/misaka4e22")
	if err != nil {
		t.Error(err)
	}
	if actor == nil {
		t.Error("Actor is nil")
	}
}

func Test_GetFollowers(t *testing.T) {
	actor, err := dbmanagers.GetActorByURI(db.Context, "https://commune1.localdomain/communities/limelight")
	if err != nil {
		t.Error(err)
	}
	if actor == nil {
		t.Error("Actor is nil")
	}

	var followers []models.Actor
	db.DB().Model(actor).Association("Followers").Find(&followers)
	found := false
	for _, follower := range followers {
		if follower.GetUsername() == "misaka4e22" {
			found = true
		}
	}

	if !found {
		t.Error("follower not found")
	}
}
func Test_GetFollowing(t *testing.T) {
	actor, err := dbmanagers.GetActorByURI(db.Context, "https://commune1.localdomain/users/misaka4e22")
	if err != nil {
		t.Error(err)
	}
	if actor == nil {
		t.Error("Actor is nil")
	}

	var followings []models.Actor
	db.DB().Model(actor).Association("Following").Find(&followings)
	found := false
	for _, following := range followings {
		if following.GetUsername() == "limelight" {
			found = true
		}
	}

	if !found {
		t.Error("following not found")
	}
}
