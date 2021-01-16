package db

import (
	"context"
	"os"
	"strings"

	"github.com/commune-project/commune/models"
	"github.com/go-redis/redis/v8"
	"github.com/rbcervilla/redisstore/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// EnvSettings describes settings in the .env
type EnvSettings struct {
	// LocalDomains contains all domains this program serves.
	LocalDomains []string
}

// SiteContext provides a common object load on launch.
type SiteContext struct {
	// Settings contains all settings in the .env.
	Settings EnvSettings
	// DB is a gorm.DB instance.
	DB    *gorm.DB
	Redis *redis.Client
	Store *redisstore.RedisStore
}

// Context is load on launch
var Context SiteContext

func init() {
	openDB()
	readSettings()
}

func openDB() {
	dbURL, isPresent := os.LookupEnv("DATABASE_URL")
	if !isPresent {
		panic("Please set DATABASE_URL environment var!")
	}
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to init db:" + err.Error())
	}
	models.ModelSetup(db)
	Context.DB = db

	redisAddr, isPresent := os.LookupEnv("REDIS_ADDR")
	if !isPresent {
		panic("Please set REDIS_ADDR environment var!")
	}

	Context.Redis = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	Context.Store, err = redisstore.NewRedisStore(context.Background(), Context.Redis)
	if err != nil {
		panic("failed to create redis store: " + err.Error())
	}

	Context.Store.KeyPrefix("session_")
}

func readSettings() {
	sLocalDomains, isPresent := os.LookupEnv("COMMUNE_LOCAL_DOMAINS")
	if !isPresent {
		Context.Settings.LocalDomains = []string{}
	} else {
		Context.Settings.LocalDomains = strings.Split(sLocalDomains, " ")
	}
}

// DB is a shortcut to db.Context.DB
func DB() *gorm.DB {
	return Context.DB
}
