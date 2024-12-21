package db

import (
	"fmt"
	"notification_service/internal/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sqlx.DB
}

func NewConnect(config *config.DbConfig) (*Repository, error) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	repository := &Repository{db: db}

	return repository, nil
}
