package util

import (
	"github.com/gin-gonic/gin"
	"slim-admin/enum"
)

type Response struct{}

func (resp *Response) Success(data any) map[string]any {
	return gin.H{
		"Success": true,
		"Data":    data,
	}
}

func (resp *Response) Failure(errCode enum.ErrorCode, errMsg string) map[string]any {
	return gin.H{
		"Success": false,
		"ErrCode": errCode,
		"ErrMsg":  errMsg,
	}
}
