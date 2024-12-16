package db

import (
	_ "embed"
	"user_service/internal/models"
)

//go:embed queries/user/createUser.sql
var createUserQuery string

//go:embed queries/user/getUser.sql
var getUserQuery string

func (r *Repository) CreateUser(phone string, passHash []byte) (string, error) {
	var userId string
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	err = tx.QueryRow(createUserQuery, phone, passHash).Scan(&userId)
	if err != nil {
		return err
	}
	tx.Commit()
	return userId, nil
}
func (r *Repository) User(phone string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Get(user, getUserQuery, phone)
	return user, err
}
