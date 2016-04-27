package controller

import (
	. "dz/web/context"

	"github.com/gin-gonic/gin"
)

func init() {
	WEB.GET("/", _home.Index)
}

type homeController struct{}

var _home = homeController{}

func (this *homeController) Index(ctx *Context) {
	ctx.HTML(200, "home.html",gin.H{"title":"首页"})
}
