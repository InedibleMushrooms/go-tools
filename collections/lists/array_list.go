package lists

import (
	"fmt"
	"sort"
)

type ArrayList[T comparable] struct {
	elements []T
	size     int
}

// NewArrayList 创建具有指定默认初始容量的 ArrayList
func NewArrayList[T comparable]() *ArrayList[T] {
	return NewArrayListWithCapacity[T](10)
}

// NewArrayListOf 根据已有的 slice 创建 ArrayList
func NewArrayListOf[T comparable](slice []T) *ArrayList[T] {
	return &ArrayList[T]{
		elements: slice,
		size:     len(slice),
	}
}

// NewArrayListWithCapacity 创建具有指定初始容量的 ArrayList
func NewArrayListWithCapacity[T comparable](initialCapacity int) *ArrayList[T] {
	if initialCapacity < 0 {
		panic("initialCapacity must be >= 0")
	}
	return &ArrayList[T]{
		elements: make([]T, initialCapacity),
		size:     0,
	}
}

// ensureCapacity 确保容量足够，不足时进行扩容
func (list *ArrayList[T]) ensureCapacity(minCapacity int) {
	oldCapacity := cap(list.elements)
	if minCapacity > oldCapacity {
		var newCapacity int
		if oldCapacity == 0 {
			// 默认初始容量
			newCapacity = 10
		} else {
			// 1.5 倍扩容
			newCapacity = oldCapacity + (oldCapacity >> 1)
		}
		if newCapacity < minCapacity {
			newCapacity = minCapacity
		}
		list.grow(newCapacity)
	}
}

// grow 执行扩容操作
func (list *ArrayList[T]) grow(newCapacity int) {
	newElements := make([]T, newCapacity)
	copy(newElements, list.elements[:list.size])
	list.elements = newElements
}

// checkIndex 索引检查
func (list *ArrayList[T]) checkIndex(index int) {
	if index < 0 || index >= list.size {
		panic(fmt.Sprintf("index %d out of bounds [0-%d]", index, list.size-1))
	}
}

// Add 添加元素到末尾
func (list *ArrayList[T]) Add(element T) {
	list.ensureCapacity(list.size + 1)
	list.elements[list.size] = element
	list.size++
}

// AddAtIndex 在指定索引插入元素
func (list *ArrayList[T]) AddAtIndex(index int, element T) {
	if index < 0 || index > list.size {
		panic("index out of bounds")
	}
	list.ensureCapacity(list.size + 1)
	// 后移元素
	copy(list.elements[index+1:], list.elements[index:list.size])
	list.elements[index] = element
	list.size++
}

// Get 获取指定索引元素
func (list *ArrayList[T]) Get(index int) T {
	list.checkIndex(index)
	return list.elements[index]
}

// Set 设置指定索引元素，返回旧值
func (list *ArrayList[T]) Set(index int, element T) T {
	list.checkIndex(index)
	old := list.elements[index]
	list.elements[index] = element
	return old
}

// Remove 移除指定索引元素，返回被移除元素
func (list *ArrayList[T]) Remove(index int) T {
	list.checkIndex(index)
	old := list.elements[index]
	// 前移元素
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	// 帮助 gc 回收
	var zero T
	list.elements[list.size] = zero
	list.elements = list.elements[:list.size]
	return old
}

// Size 获取当前元素数量
func (list *ArrayList[T]) Size() int {
	return list.size
}

// IsEmpty 判断是否为空
func (list *ArrayList[T]) IsEmpty() bool {
	return list.size == 0
}

// Contains 判断是否包含元素
func (list *ArrayList[T]) Contains(element T) bool {
	for i := 0; i < list.size; i++ {
		if list.elements[i] == element {
			return true
		}
	}
	return false
}

// IndexOf 查找元素首次出现的索引
func (list *ArrayList[T]) IndexOf(element T) int {
	for i := 0; i < list.size; i++ {
		if list.elements[i] == element {
			return i
		}
	}
	return -1
}

// Clear 清空列表
func (list *ArrayList[T]) Clear() {
	list.elements = make([]T, 0)
	list.size = 0
}

// Iterator 迭代器支持 todo 抽象为接口
func (list *ArrayList[T]) Iterator() <-chan T {
	ch := make(chan T)
	go func() {
		for i := 0; i < list.size; i++ {
			ch <- list.elements[i]
		}
		close(ch)
	}()
	return ch
}

// Sort 排序支持 todo 抽象为接口
func (list *ArrayList[T]) Sort(less func(a, b T) bool) {
	sort.Slice(list.elements[:list.size], func(i, j int) bool {
		return less(list.elements[i], list.elements[j])
	})
}
