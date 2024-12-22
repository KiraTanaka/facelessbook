package db

import (
	_ "embed"
	"fmt"
	"user_service/internal/models"
)

//go:embed queries/user/createUser.sql
var createUserQuery string

//go:embed queries/user/getUser.sql
var getUserQuery string

//go:embed queries/user/getNickname.sql
var getNicknameQuery string

func (r *Repository) CreateUser(phone string, passHash []byte) (string, error) {
	var userId string
	tx, err := r.db.Beginx()
	if err != nil {
		return "", fmt.Errorf("begin a transaction for create user: %w", err)
	}
	defer tx.Rollback()
	err = tx.QueryRow(createUserQuery, phone, passHash).Scan(&userId)
	if err != nil {
		return "", fmt.Errorf("create user: %w", err)
	}
	tx.Commit()
	return userId, nil
}
func (r *Repository) User(phone string) (*models.User, error) {
	user := &models.User{}
	if err := r.db.Get(user, getUserQuery, phone); err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}
	return user, nil
}

func (r *Repository) Nickname(userId string) (string, error) {
	var nick string
	if err := r.db.Get(&nick, getNicknameQuery, userId); err != nil {
		return "", fmt.Errorf("get nickname: %w", err)
	}
	return nick, nil
}
