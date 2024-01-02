package core

// Result 结果集接口
type Result interface {
	Error
}

// MapResult 数据结果集
type MapResult[K DRKey, V any] interface {

	// Get 根据key获取值
	//
	// @param key 键
	Get(key K) V

	// Put 添加值
	//
	// @param key 键
	//
	// @param value 值
	Put(key K, value V)

	// PutIfAbsent 添加值
	//
	// @param key 键
	//
	// @param value 值
	PutIfAbsent(key K, value V)

	// PutAll 添加多个值
	//
	// @param data map值
	PutAll(data map[K]V)

	// PutAllIfAbsent 添加多个值
	//
	// @param data map值
	PutAllIfAbsent(data map[K]V)

	// Remove 根据指定键移除值
	//
	// @param key 键
	Remove(key K)

	// Clear 清空数据
	Clear()

	// ContainsKey 检查是否包含指定的键
	//
	// @param key 键
	ContainsKey(key K) bool

	// ContainsValue 检查是否包含指定的值
	//
	// @param value 值
	ContainsValue(value V) bool

	// GetSize 获取元素个数
	GetSize() int

	// IsEmpty 检查是否存在数据
	IsEmpty() bool
}
