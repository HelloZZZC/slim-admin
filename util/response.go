package util

import "github.com/gin-gonic/gin"

type Response struct{}

func (resp *Response) Success(data any) map[string]any {
	return gin.H{
		"success": true,
		"data":    data,
	}
}

func (resp *Response) Failure(errCode string, errMsg string) map[string]any {
	return gin.H{
		"success": false,
		"errCode": errCode,
		"errMsg":  errMsg,
	}
}
