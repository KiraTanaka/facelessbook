package models

type User struct {
	Id       string `json:"id" db:"id" binding:"max=36"`
	Phone    string `json:"phone" db:"phone" binding:"required,max=11"`
	PassHash []byte `json:"pass_hash" db:"pass_hash" binding:"required"`
}
