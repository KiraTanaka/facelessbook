package repository

import (
	"facelessbook/post_service/app/internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"

	log "github.com/sirupsen/logrus"
)

type repository struct {
	db *sqlx.DB
}

func NewDbConnect(config *config.Config) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Error(err)
	}

	err = db.Ping()
	if err != nil {
		log.Error(err)
	}
}
