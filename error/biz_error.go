package error

import (
	"errors"
	"fmt"
	"slim-admin/enum"
)

type BizError struct {
	Code enum.ErrorCode
	Msg  string
}

func (e BizError) Error() string {
	return fmt.Sprintf("Biz Error: code: [%s]，Msg: %s", e.Code, e.Msg)
}

// NewBizError 初始化业务异常
func NewBizError(code enum.ErrorCode, msg string) BizError {
	return BizError{Code: code, Msg: msg}
}

// Transform 异常转换
func Transform(err error) BizError {
	if !errors.As(err, &BizError{}) {
		return NewBizError(enum.UnknownSystemError, "未知系统错误，请联系管理员")
	}

	return err.(BizError)
}
