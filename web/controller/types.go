package controller

import (
	"dz/service/web"
	. "dz/web/context"
	"github.com/gin-gonic/gin"
)

var (
	_v1 = "1.0"
	V1  = &APIVersion{
		Version:     _v1,
		RouterGroup: web.App.Group(_v1),
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
