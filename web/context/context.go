package context

import (
	"dz/dal/context"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	DBCtx *context.Context
}

func NewContext(*gin.Context) *Context {
	return &Context{}
}
