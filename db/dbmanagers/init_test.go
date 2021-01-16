package dbmanagers_test

import (
	"os"
	"testing"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
)

func init() {
	models.DeleteFromTables(db.Context.DB)
	db.Seeding()
}

func deinit() {
	models.DeleteFromTables(db.Context.DB)
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	code := m.Run()
	deinit()
	os.Exit(code)
}
