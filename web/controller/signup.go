package controller

import (
	. "dz/web/context"

	"github.com/gin-gonic/gin"
//    "html/template"
)

func init() {
	WEB.GET("/xsm", _signup.XSM)
//	WEB.POST("/xsm", _signup.XSMSubmit)
	WEB.GET("/qpg", _signup.QPG)
//	WEB.POST("/qpg", _signup.QPGSubmit)
	WEB.GET("/jc", _signup.JC)
	WEB.POST("/jc", _signup.JCSubmit)
}

type signupController struct{}

var _signup = signupController{}

func (this *signupController) XSM(ctx *Context) {
	ctx.HTML(200, "xsm_signup.html",gin.H{"title":"小树苗报名"})
}

func (this *signupController) QPG(ctx *Context) {
	ctx.HTML(200, "qpg_signup.html",gin.H{"title":"青苹果报名"})
}

func (this *signupController) JC(ctx *Context) {
	ctx.HTML(200, "jc_signup.html",gin.H{"title":"南京初阶二日禅报名表4/30-5/1"})
}

func (this *signupController) JCSubmit(ctx *Context) {
    ctx.Data(200,"text/html",[]byte(`<a href="/1.0/jc">返回</a>`))
}
