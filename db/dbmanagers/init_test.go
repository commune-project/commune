package dbmanagers_test

import (
	"os"
	"testing"

	"github.com/commune-project/commune/db"
)

func init() {
	db.Seeding()
}

func deinit() {

}

func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	code := m.Run()
	deinit()
	os.Exit(code)
}
