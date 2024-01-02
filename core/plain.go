package core

// PlainResult 单值结果集
type PlainResult[T any] struct {
	*Brief
	// Data 数据
	Data T `json:"data"`
}

func (p *PlainResult[T]) GetData() T {
	return p.Data
}

func (p *PlainResult[T]) SetData(data T) {
	p.Data = data
}

// NewPlainWithEmpty 创建实例
func NewPlainWithEmpty[T any]() *PlainResult[T] {
	_result := &PlainResult[T]{
		Brief: &Brief{
			Status:  200,
			Code:    200,
			Success: true,
		},
	}
	_result._setTime()
	return _result
}

// NewPlainWithData 创建实例
//
// @param data 数据
func NewPlainWithData[T any](data T) *PlainResult[T] {
	_result := &PlainResult[T]{
		Data: data,
		Brief: &Brief{
			Status:  200,
			Code:    200,
			Success: true,
		},
	}
	_result._setTime()
	return _result
}

// NewPlainFailure 创建实例
//
// @param status HTTP状态值
//
// @param code 状态码
//
// @param msg 异常信息
func NewPlainFailure[T any](status int32, code int32, msg string) *PlainResult[T] {
	_result := &PlainResult[T]{
		Brief: &Brief{
			Status:  status,
			Code:    code,
			Msg:     msg,
			Success: false,
		},
	}
	_result._setTime()
	return _result
}

// NewPlainFailureWithEmpty 创建实例
func NewPlainFailureWithEmpty[T any]() *PlainResult[T] {
	return NewPlainFailure[T](200, 500500, "Server Error")
}

// NewPlainFailureWithMsg 创建实例
//
// @param code 状态码
//
// @param msg 异常信息
func NewPlainFailureWithMsg[T any](code int32, msg string) *PlainResult[T] {
	return NewPlainFailure[T](200, code, msg)
}
