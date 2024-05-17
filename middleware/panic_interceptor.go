package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_error "slim-admin/error"
	"slim-admin/util"
)

func PanicInterceptor() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				bizError := _error.Transform(r.(error))
				ctx.JSON(http.StatusOK, new(util.Response).Failure(bizError.Code, bizError.Msg))
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
