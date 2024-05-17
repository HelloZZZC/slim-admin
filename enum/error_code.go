package enum

type ErrorCode string

const (
	UnknownSystemError ErrorCode = "-1"       // 未知系统异常
	ExampleError       ErrorCode = "01010001" // 业务异常示例
)
