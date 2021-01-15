package communectl

import (
	"flag"
	"fmt"
	"os"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/misaka4e21/gorm-migrator/migrator"
)

func helpSubCmd(subcmd string) {
	switch subcmd {
	case "migrate":
		fmt.Println(`Usage:
		commune migrate up [<steps>]
		commune migrate down [<steps>]
		commune migrate force <version>
		commune migrate version`)
		os.Exit(0)
	default:
		fmt.Println("Supported subcommands: generate-migrations, migrate")
		os.Exit(0)
	}
}

// Communectl deals with subcommands.
func Communectl() {
	if len(os.Args) < 2 {
		helpSubCmd("")
	} else if os.Args[1] == "generate-migrations" {
		// variables declaration
		var migrationName string

		// flags declaration using flag package
		flag.CommandLine.StringVar(&migrationName, "name", "migration", "Specify migration's name. Default is migration.")

		flag.CommandLine.Parse(os.Args[2:]) // after declaring flags we need to call it
		err := migrator.GenerateMigrations(migrationName, db.DB, &models.Account{}, &models.User{}, &models.CommuneMember{}, &models.LocalCommune{}, &models.Category{}, &models.Commune{}, &models.Post{})
		if err != nil {
			panic(err.Error())
		}
	} else if os.Args[1] == "migrate" {
		migrateSubCmd()
	} else if os.Args[1] == "help" {
		if len(os.Args) >= 3 {
			helpSubCmd(os.Args[2])
		}
	} else {
		helpSubCmd("")
	}
}
