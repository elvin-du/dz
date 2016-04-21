package model

type User struct {
	Id   string `db:id`
	Name string `db:name`
}

type Post struct {
	Id      uint64 `db:id`
	Title   string `db:title`
	Content string `db:content`
}
