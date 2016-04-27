package controller

import (
	"dz/service/web"
    . "dz/web/controller"
)

var (
	_admin = "admin"
	V1Admin = &APIVersion{
		Version:     _admin,
		RouterGroup: web.App.Group(_admin),
	}
)
