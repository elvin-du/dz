package context

import (
	"database/sql"
	"dz/service/db"
//	"dz/service/log"
)

type Context struct {
	DB *sql.DB
}

func NewContext() *Context {
	return &Context{DB: db.DB()}
}

func (this *Context) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return this.DB.Query(query, args...)
}
