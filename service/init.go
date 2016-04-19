package service

import (
	"dz/service/config"
	"dz/service/db"
	dzlog "dz/service/log"
	"dz/service/web"
)

func Init() {
	config.Init()
	dzlog.Init()
	db.Init()
	web.Init()
}
