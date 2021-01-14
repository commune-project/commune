package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/commune-project/commune/db"
	"github.com/commune-project/commune/models"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/misaka4e21/gorm-migrator/migrator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Supported subcommands: generate-migrations")
	} else if os.Args[1] == "generate-migrations" {
		// variables declaration
		var migrationName string

		// flags declaration using flag package
		flag.CommandLine.StringVar(&migrationName, "n", "migration", "Specify migration's name. Default is migration.")

		flag.CommandLine.Parse(os.Args[2:]) // after declaring flags we need to call it
		err := migrator.GenerateMigrations(migrationName, db.DB, &models.Account{}, &models.User{}, &models.CommuneMember{}, &models.LocalCommune{}, &models.Category{}, &models.Commune{}, &models.Post{})
		if err != nil {
			panic(err.Error())
		}
	} else if os.Args[1] == "migrate" {
		db, err := db.DB.DB()
		if err != nil {
			panic(err)
		}
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		m, err := migrate.NewWithDatabaseInstance(
			"file://migrations",
			"postgres", driver)
		if err != nil {
			panic(err.Error())
		}
		m.Steps(2)
	}
}
