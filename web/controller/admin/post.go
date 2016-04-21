package controller

import (
	. "dz/web/context"

	"dz/dw"

	"dz/dal/model"
	"dz/service/log"

	"github.com/gin-gonic/gin"
)

func init() {
	V1Admin.GET("/post", _post.Index)
	V1Admin.POST("/post", _post.Do)
}

type postController struct{}

var _post = postController{}

func (this *postController) Index(ctx *Context) {
	ctx.HTML(200, "admin_post.html", gin.H{"title": "发布文章"})
}

func (this *postController) Do(ctx *Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	kind := ctx.PostForm("kind")

	post, err := model.PostModel().Add(ctx.DBCtx, title, content)
	if nil != err {
		log.Errorln(err)
		ctx.String(500, err.Error())
		return
	}

	dw.Add(post.Id, title+content+kind)
	ctx.String(200, "add ok")
}
