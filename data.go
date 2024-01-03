package next_result

// Data 数据模型接口
type Data[T any] interface {

	// GetData 获取数据
	GetData() T

	// SetData 设置数据
	//
	// @param data 数据
	SetData(data T)
}
