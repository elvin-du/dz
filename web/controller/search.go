package controller

import (
	"dz/dal/model"
	"dz/dw"
	. "dz/web/context"

	"dz/service/log"
	"strings"

	"github.com/gin-gonic/gin"
    "html/template"
)

func init() {
	V1.GET("/search", _search.Index)
	V1.POST("/search", _search.Do)
}

type searchController struct{}

var _search = searchController{}

func (this *searchController) Index(ctx *Context) {
	ctx.HTML(200, "search.html", gin.H{"title": "搜索"})
}

func (this *searchController) Do(ctx *Context) {
	key := ctx.PostForm("key")
	resp := dw.Query(key)
	log.Debugln(resp)
	post, err := model.PostModel().Query(ctx.DBCtx, resp.Docs[0].DocId)
	if nil != err {
		log.Errorln(err)
		ctx.String(200, err.Error())
		return
	}

    log.Debugln(post.Content)
	res := strings.Replace(post.Title+post.Content, resp.Tokens[0], "<strong>"+resp.Tokens[0]+"</strong>", -1)
	ctx.HTML(200, "search_result.html", gin.H{"title": "搜索结果", "result": template.HTML(res)})
}
