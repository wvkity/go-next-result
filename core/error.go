package core

// Error Error接口
type Error interface {

	// GetStatus 获取HTTP状态码
	GetStatus() int32

	// SetStatus 设置HTTP状态码
	//
	// @param status HTTP状态码
	SetStatus(status int32)

	// GetCode 获取状态码
	GetCode() int32

	// SetCode 设置状态码
	//
	// @param code 状态码
	SetCode(code int32)

	// GetMsg 获取异常信息
	GetMsg() string

	// SetMsg 设置异常信息
	//
	// @param msg 异常信息
	SetMsg(msg string)
}
