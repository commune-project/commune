package dbmanagers_test

import (
	"testing"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
)

func Test_GetAccountByUsername(t *testing.T) {
	account, err := dbmanagers.GetAccountByUsername(db.DB, "misaka4e22", "commune1.localdomain")
	if err != nil {
		t.Error(err)
	}
	if account == nil {
		t.Error("account is nil")
	}
}

func Test_GetAccountByUsernameWithDifferentDomain(t *testing.T) {
	account, err := dbmanagers.GetAccountByUsername(db.DB, "misaka4e22", "not.exist.localdomain")
	if err == nil && account != nil {
		t.Error(account.GetUsername() + "@" + account.GetDomain() + "shouldn't exist")
	}
}
