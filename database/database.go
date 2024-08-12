package database

import (
	"fmt"
	"sync"

	"github.com/eslami200117/clientCli/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *PostgresDatabase
)

func NewPostgresDatabase(conf *config.Config) *PostgresDatabase {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			conf.Db.Host,
			conf.Db.User,
			conf.Db.Password,
			conf.Db.DBName,
			conf.Db.Port,
			conf.Db.SSLMode,
			conf.Db.TimeZone,
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		dbInstance = &PostgresDatabase{Db: db}
	})

	return dbInstance
}

func (p *PostgresDatabase) GetDb() *gorm.DB {
	return p.Db
}
