package web

import (
	dzlog "dz/service/log"

	"github.com/gin-gonic/gin"
)

const (
	HTTP_ADDR = `:40001`
)

var (
	app *gin.Engine
)

func Init() {
	app = gin.New()
	dzlog.Infof("Runing on:%s", HTTP_ADDR)

	err := app.Run(HTTP_ADDR)
	if nil != err {
		dzlog.Fatalln(err)
	}
}

func Web() *gin.Engine {
	return app
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

func InitRouters() {

}

type RouterGroup *gin.RouterGroup
