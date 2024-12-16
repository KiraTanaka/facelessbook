package models

type User struct {
	Id       string
	Phone    string
	PassHash []byte
}
