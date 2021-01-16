package communectl

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/commune-project/commune/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // Migration source
	_ "github.com/lib/pq"
)

func migrateSubCmd() {
	db, err := db.DB().DB()
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

	step := -1
	if len(os.Args) >= 4 {
		step, err = strconv.Atoi(os.Args[3])

		if err != nil {
			helpSubCmd("migrate")
		}
	}

	if len(os.Args) < 3 {
		helpSubCmd("migrate")
	}

	subcommand := os.Args[2]

	// handle Ctrl+c
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT)
	go func() {
		for range signals {
			log.Println("Stopping after this running migration ...")
			m.GracefulStop <- true
			return
		}
	}()

	switch subcommand {
	case "force":
		if step != -1 {
			m.Force(step)
		} else {
			helpSubCmd("migrate")
		}
	case "version":
		if step == -1 {
			version, dirty, err := m.Version()
			if err != nil {
				panic("Error getting version.")
			}
			fmt.Println(version)
			if dirty {
				fmt.Println("Database schema is dirty.")
			}
		} else {
			helpSubCmd("migrate")
		}
	case "up":
		err = upCmd(m, step)
	case "down":
		err = downCmd(m, step)
	default:
		helpSubCmd("migrate")
	}
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	os.Exit(0)
}

func upCmd(m *migrate.Migrate, limit int) error {
	if limit >= 0 {
		if err := m.Steps(limit); err != nil {
			if err != migrate.ErrNoChange {
				return err
			}
			log.Println(err)
		}
	} else {
		if err := m.Up(); err != nil {
			if err != migrate.ErrNoChange {
				return err
			}
			log.Println(err)
		}
	}
	return nil
}

func downCmd(m *migrate.Migrate, limit int) error {
	if limit >= 0 {
		if err := m.Steps(-limit); err != nil {
			if err != migrate.ErrNoChange {
				return err
			}
			log.Println(err)
		}
	} else {
		if err := m.Down(); err != nil {
			if err != migrate.ErrNoChange {
				return err
			}
			log.Println(err)
		}
	}
	return nil
}
