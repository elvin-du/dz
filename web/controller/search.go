package controller

import (
	"dz/dal/model"
	"dz/dw"
	. "dz/web/context"

	"dz/service/log"
	"strings"

	"html/template"

	"github.com/gin-gonic/gin"
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

	var boo = []map[string]template.HTML{}
	for _, v := range resp.Docs {
		post, err := model.PostModel().Query(ctx.DBCtx, v.DocId)
		if nil != err {
			log.Errorln(err)
			ctx.String(200, err.Error())
			return
		}
        log.Debugln(v.TokenSnippetLocations)
        log.Debugln(v.TokenLocations)
        log.Debugln(v.Scores)

		title := post.Title
		content := post.Content
		for _, v2 := range resp.Tokens {
            log.Debugln(v2)
			title = strings.Replace(title, v2, `<span style="color:red">`+v2+`</span>`, -1)
//            start := v.TokenSnippetLocations[i]
//            if start < 30{
//                start = 0
//            }else{
//                start = -30
//            }

//            end := len(content) - v.TokenSnippetLocations[i]
//            if end >= 30{
//                end = v.TokenSnippetLocations[i] + 30
//            }else{
//                end = len(content) - v.TokenSnippetLocations[i]
//            }

			content = strings.Replace(content, v2, `<span style="color:red">`+v2+`</span>`, -1)
		}

		var m = map[string]template.HTML{"title": template.HTML(title), "content": template.HTML(content)}
		boo = append(boo, m)
	}

	ctx.HTML(200, "search_result.html", gin.H{"title": "搜索结果", "results": boo})
}
