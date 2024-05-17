package api

import (
	"github.com/gin-gonic/gin"
	"slim-admin/enum"
	_error "slim-admin/error"
)

type ExampleController struct{}

func (c *ExampleController) BizError(ctx *gin.Context) {
	err := _error.NewBizError(enum.ExampleError, "异常示例")
	panic(err)
}
