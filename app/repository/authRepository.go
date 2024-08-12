package repository

import (
	"github.com/eslami200117/clientCli/app/entities"
	"github.com/eslami200117/clientCli/database"
)

type repository struct {
	db *database.PostgresDatabase
}


func NewRepo(db *database.PostgresDatabase) *repository{
	return &repository{
		db: db,
	}
}


func (pr repository) InsertAuth(username string, token string) {
	var auth entities.AuthEntity
	pr.db.GetDb().First(&auth, "username= ?", username)
	if auth.Username == username {
		auth.AuthToken = token
		pr.db.GetDb().Save(&auth)
	} else {
		pr.db.GetDb().Create(
			&entities.AuthEntity{
				Username: username,
				AuthToken: token,
			},
		)
	}
}