package model

type User struct {
	Id   string `db:id`
	Name string `db:name`
}
