package web

import (
	dzlog "dz/service/log"

	"github.com/gin-gonic/gin"
)

const (
	HTTP_ADDR = `:40001`
)

var (
	App = gin.New()
)

func Start() {
	dzlog.Infof("Runing on:%s", HTTP_ADDR)
    App.LoadHTMLGlob("static/tpl/*")
    App.Static("/pub","./static")
	err := App.Run(HTTP_ADDR)
	if nil != err {
		dzlog.Fatalln(err)
	}
}

type HttpHandler func(*gin.Context)

var (
	handlers = map[string]HttpHandler{}
)

func Register(path string, h HttpHandler) {
	if _, ok := handlers[path]; ok {
		dzlog.Fatalf("router:%s already registered!", path)
	}

	handlers[path] = h
}
