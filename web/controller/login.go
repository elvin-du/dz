package controller

import (
	. "dz/web/context"

	"github.com/gin-gonic/gin"
)

func init() {
	WEB.GET("/login", _login.Index)
}

type loginController struct{}

var _login = loginController{}

func (this *loginController) Index(ctx *Context) {
	ctx.HTML(200, "login.html",gin.H{"title":"登录"})
}
