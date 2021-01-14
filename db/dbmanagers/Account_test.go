package dbmanagers_test

import (
	"testing"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/db/dbmanagers"
)

func Test_GetAccountByUsername(t *testing.T) {
	account, err := dbmanagers.GetAccountByUsername(db.DB, "misaka4e22")
	if err != nil {
		t.Error(err)
	}
	if account == nil {
		t.Error("account is nil")
	}
}
