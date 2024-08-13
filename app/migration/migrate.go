package main

import (
	"github.com/eslami200117/clientCli/app/entities"
	"github.com/eslami200117/clientCli/config"
	"github.com/eslami200117/clientCli/database"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	authMigrate(db)
}

func authMigrate(db *database.PostgresDatabase) {
	db.GetDb().Migrator().CreateTable(&entities.AuthEntity{})

}
