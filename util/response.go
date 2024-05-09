package util

import "github.com/gin-gonic/gin"

type Response struct{}

func (r *Response) Success(data any) map[string]any {
	return gin.H{
		"success": true,
		"data":    data,
	}
}
