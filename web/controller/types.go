package controller

import (
	"dz/service/web"
	. "dz/web/context"
	"github.com/gin-gonic/gin"
)

var (
	WEB  = &APIVersion{
		RouterGroup: web.App.Group(""),
	}
)

type APIVersion struct {
	Version     string
	RouterGroup *gin.RouterGroup
}

type Handler func(*Context)

func (this *APIVersion) GET(relativePath string, handlers ...Handler) {
	this.RouterGroup.GET(relativePath, func(ctx *gin.Context) {
		for _, f := range handlers {
			f(NewContext(ctx))
		}
	})
}

func (this *APIVersion) POST(relativePath string, handlers ...Handler) {
	this.RouterGroup.POST(relativePath, func(ctx *gin.Context) {
		for _, f := range handlers {
			f(NewContext(ctx))
		}
	})
}
