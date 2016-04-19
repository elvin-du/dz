package controller

import (
	"dz/service/web"

	. "dz/web/context"

	"github.com/gin-gonic/gin"
)

const (
	v1 string = "1.0"
)

func V1() *gin.RouterGroup {
	return web.Web().Group(v1)
}

type Handler func(*Context)

func GET(relativePath string, handlers ...Handler) {
	V1().GET(relativePath, func(ctx *gin.Context) {
		for _, f := range handlers {
			f(NewContext(ctx))
		}
	})
}
