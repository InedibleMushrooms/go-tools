package interfaces

import "reflect"

type Object interface {

	// Equals 判断是否相等
	Equals(other interface{}) bool

	// HashCode hash code 值
	HashCode() int

	// ToString string 值
	ToString() string

	// Class // 获取类型信息
	Class() reflect.Type

	// Clone 浅克隆方法
	Clone() (Object, error)
}
