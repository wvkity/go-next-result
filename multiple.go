package next_result

import (
	"reflect"
)

// MultipleResult 多值结果集
type MultipleResult[K DRKey, V any] struct {
	*Brief
	// Map数据
	Data map[K]V `json:"data"`
}

func (m *MultipleResult[K, V]) GetData() map[K]V {
	return m.Data
}

func (m *MultipleResult[K, V]) SetData(data map[K]V) {
	if data == nil {
		m.Data = make(map[K]V, 8)
	} else {
		m.Data = data
	}
}

func (m *MultipleResult[K, V]) Get(key K) V {
	return m.Data[key]
}

func (m *MultipleResult[K, V]) Put(key K, value V) {
	m.Data[key] = value
}

func (m *MultipleResult[K, V]) PutIfAbsent(key K, value V) {
	if !m.ContainsKey(key) {
		m.Data[key] = value
	}
}

func (m *MultipleResult[K, V]) PutAll(data map[K]V) {
	if data != nil && len(data) > 0 {
		for key, value := range data {
			m.Data[key] = value
		}
	}
}

func (m *MultipleResult[K, V]) PutAllIfAbsent(data map[K]V) {
	if data != nil && len(data) > 0 {
		for key, value := range data {
			if !m.ContainsKey(key) {
				m.Data[key] = value
			}
		}
	}
}

func (m *MultipleResult[K, V]) Remove(key K) {
	if m.ContainsKey(key) {
		delete(m.Data, key)
	}
}

func (m *MultipleResult[K, V]) Clear() {
	m.Data = make(map[K]V, 8)
}

func (m *MultipleResult[K, V]) ContainsKey(key K) bool {
	if m.Data == nil {
		return false
	}
	_, hasKey := m.Data[key]
	return hasKey
}

func (m *MultipleResult[K, V]) ContainsValue(value V) bool {
	if !m.IsEmpty() {
		for _, v := range m.Data {
			if reflect.DeepEqual(v, value) {
				return true
			}
		}
	}
	return false
}

func (m *MultipleResult[K, V]) GetSize() int {
	if m.Data == nil {
		return 0
	}
	return len(m.Data)
}

func (m *MultipleResult[K, V]) IsEmpty() bool {
	return m.Data == nil || len(m.Data) == 0
}

// NewMultipleWithEmpty 创建实例
func NewMultipleWithEmpty[K DRKey, V any]() *MultipleResult[K, V] {
	_result := &MultipleResult[K, V]{
		Data: map[K]V{},
		Brief: &Brief{
			Status:  200,
			Code:    200,
			Success: true,
		},
	}
	_result._setTime()
	return _result
}

// NewMultipleWithSize 创建实例
//
// @param size 容量大小
func NewMultipleWithSize[K DRKey, V any](size int) *MultipleResult[K, V] {
	_result := &MultipleResult[K, V]{
		Data: make(map[K]V, size),
		Brief: &Brief{
			Status:  200,
			Code:    200,
			Success: true,
		},
	}
	_result._setTime()
	return _result
}

// NewMultipleWithOne 创建实例
//
// @param key 键
//
// @param value 值
func NewMultipleWithOne[K DRKey, V any](key K, value V) *MultipleResult[K, V] {
	_result := NewMultipleWithSize[K, V](8)
	_result.Put(key, value)
	return _result
}

// NewMultipleWithMany 创建实例
//
// @param data 多个值
func NewMultipleWithMany[K DRKey, V any](data map[K]V) *MultipleResult[K, V] {
	var _size int
	if data == nil || len(data) == 0 {
		_size = 8
	} else {
		_size = len(data)
	}
	_result := NewMultipleWithSize[K, V](_size)
	_result.PutAll(data)
	return _result
}

// NewMultipleFailure 创建实例
//
// @param status HTTP状态码
//
// @param code 状态码
//
// @param msg 错误消息
func NewMultipleFailure[K DRKey, V any](status int32, code int32, msg string) *MultipleResult[K, V] {
	_result := &MultipleResult[K, V]{
		Data: map[K]V{},
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

// NewMultipleFailureWithEmpty 创建实例
func NewMultipleFailureWithEmpty[K DRKey, V any]() *MultipleResult[K, V] {
	return NewMultipleFailure[K, V](200, 500500, "Server Error")
}

// NewMultipleFailureWithMsg 创建实例
//
// @param code 状态码
//
// @param msg 错误信息
func NewMultipleFailureWithMsg[K DRKey, V any](code int32, msg string) *MultipleResult[K, V] {
	return NewMultipleFailure[K, V](200, code, msg)
}
