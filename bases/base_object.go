package bases

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"go-tools/bases/interfaces"
	"reflect"
)

// BaseObject 基础实现（可被其他结构体组合）
type BaseObject struct{}

// Equals 默认相等性判断（按地址比较）
func (b *BaseObject) Equals(other interface{}) bool {
	return reflect.DeepEqual(b, other)
}

// HashCode 默认哈希实现（基于内存地址）
func (b *BaseObject) HashCode() int {
	ptrVal := reflect.ValueOf(b).Pointer()
	return int(ptrVal ^ (ptrVal >> 32))
}

// ToString 默认字符串表示
func (b *BaseObject) ToString() string {
	return fmt.Sprintf("%T@%x", b, b.HashCode())
}

// Class 获取类型信息
func (b *BaseObject) Class() reflect.Type {
	return reflect.TypeOf(b)
}

// Clone 浅克隆
func (b *BaseObject) Clone() (interfaces.Object, error) {
	// 利用反射创建新实例
	newObj := reflect.New(reflect.TypeOf(b).Elem())
	srcValue := reflect.ValueOf(b).Elem()
	newObj.Elem().Set(srcValue)
	return newObj.Interface().(interfaces.Object), nil
}

// DeepClone 深克隆
func DeepClone(obj interfaces.Object) (interfaces.Object, error) {
	// 序列化方案（推荐）
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)

	if err := enc.Encode(obj); err != nil {
		return nil, err
	}

	copyObject := reflect.New(reflect.TypeOf(obj).Elem()).Interface().(interfaces.Object)
	if err := dec.Decode(copyObject); err != nil {
		return nil, err
	}
	return copyObject, nil
}
