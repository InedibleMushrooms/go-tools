package interfaces

/*
List 接口
该接口只定义清楚各个方法的行为和表现
*/
type List[T comparable] interface {
	// Add 添加元素
	Add(element T)

	// AddAtIndex 指定位置添加元素
	AddAtIndex(index int, element T)

	// Get 获取指定位置元素
	Get(index int) T

	// Set 设置指定位置元素
	Set(index int, element T) T

	// Remove 移除指定位置元素
	Remove(index int) T

	// Size list 大小
	Size() int

	// IsEmpty list 是否为空
	IsEmpty() bool

	// Contains list 是否包含某个元素
	Contains(element T) bool

	// IndexOf 某个元素对象的下标
	IndexOf(element T) int

	// Clear 清空 list
	Clear()
}
