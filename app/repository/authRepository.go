package repository

import (
	"github.com/eslami200117/clientCli/app/entities"
	"github.com/eslami200117/clientCli/database"
)

type Repository struct {
	db *database.PostgresDatabase
}

func NewRepo(db *database.PostgresDatabase) *Repository {
	return &Repository{
		db: db,
	}
}

func (pr *Repository) InsertAuth(username string, token string) {
	var auth entities.AuthEntity
	pr.db.GetDb().Find(&auth, "username= ?", username)
	if auth.Username == username {
		auth.AuthToken = token
		pr.db.GetDb().Save(&auth)
	} else {
		pr.db.GetDb().Create(
			&entities.AuthEntity{
				Username:  username,
				AuthToken: token,
			},
		)
	}
}

func (pr *Repository) GetToken(username string) string {
	var auth entities.AuthEntity
	pr.db.GetDb().Find(&auth, "username= ?", username)
	if auth.Username != "" {
		return auth.AuthToken
	}
	return ""
}

func (pr *Repository) Logout(username string) {
	var auth entities.AuthEntity
	pr.db.GetDb().Find(&auth, "username= ?", username)
	if auth.Username != "" {
		auth.AuthToken = ""
	}
}
