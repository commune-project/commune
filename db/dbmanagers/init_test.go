package dbmanagers_test

import (
	"os"
	"testing"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
)

func init() {
	models.Migrate(db.DB)
}

func deinit() {
	models.DropTables(db.DB)
}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	code := m.Run()
	deinit()
	os.Exit(code)
}
