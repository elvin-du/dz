package service

import (
	dzlog "dz/service/log"

	"github.com/gin-gonic/gin"
)

const (
	HTTP_ADDR = `:40001`
)

func Init() {
	dzlog.Init()

	app := gin.New()
	dzlog.Infof("Runing on:%s", HTTP_ADDR)
	err := app.Run(HTTP_ADDR)
	if nil != err {
		dzlog.Fatalln(err)
	}
}
