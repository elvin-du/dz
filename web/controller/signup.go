package controller

import (
	. "dz/web/context"

	"github.com/gin-gonic/gin"
//    "html/template"
)

func init() {
	V1.GET("/xsm", _signup.XSM)
//	V1.POST("/xsm", _signup.XSMSubmit)
	V1.GET("/qpg", _signup.QPG)
//	V1.POST("/qpg", _signup.QPGSubmit)
	V1.GET("/jc", _signup.JC)
	V1.POST("/jc", _signup.JCSubmit)
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
