package controller

import (
	. "dz/web/context"
)

func init() {
	GET("/", _auth.Auth)
}

type authController struct{}

var _auth = authController{}

func (this *authController) Auth(ctx *Context) {

}
