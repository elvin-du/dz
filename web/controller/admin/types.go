package controller

import (
	"dz/service/web"
    . "dz/web/controller"
)

var (
	_v1    = "1.0"
	_admin = "admin"
	V1Admin = &APIVersion{
		Version:     _v1,
		RouterGroup: web.App.Group(_v1).Group(_admin),
	}
)
