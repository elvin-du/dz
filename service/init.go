package service

import (
	"dz/service/config"
	"dz/service/db"
	"dz/service/log"
	"dz/service/web"
	_ "dz/web/controller"
	_ "dz/web/controller/admin"
)

func Init() {
	config.Init()
	log.Init()
	db.Init()
	web.Start()
}
