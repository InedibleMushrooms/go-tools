package objects

import (
	"fmt"
	"testing"
)

type Person struct {
	BaseObject
	Name string
	Tags []string
}

// 重写 ToString
func (p *Person) ToString() string {
	return fmt.Sprintf("Person{Name=%s, Tags=%v}", p.Name, p.Tags)
}

var (
	obj1 = &Person{
		Name: "Alice",
		Tags: []string{"tag1", "tag2"},
	}
	obj2 = &Person{
		Name: "Bob",
		Tags: []string{"tag3", "tag4"},
	}
)

func TestBaseObject_Equals(t *testing.T) {
	t.Log(obj1.Name)
	t.Log(obj2.Name)
	if obj1.Equals(obj2) {
		t.Error("obj1 and obj2 should not be equal")
	}
}

func TestBaseObject_HashCode(t *testing.T) {
	t.Log(obj1.HashCode())
	t.Log(obj2.HashCode())
	if obj1.HashCode() == obj2.HashCode() {
		t.Error("obj1 and obj2 should not be equal")
	}
}

func TestBaseObject_ToString(t *testing.T) {
	t.Log(obj1.ToString())
	t.Log(obj2.ToString())
	if obj1.ToString() == obj2.ToString() {
		t.Error("obj1 and obj2 should not be equal")
	}
}

func TestBaseObject_Class(t *testing.T) {
	t.Log(obj1.Class())
	t.Log(obj2.Class())
	if obj1.Class() != obj2.Class() {
		t.Error("obj1 and obj2 should be equal")
	}
}

func TestBaseObject_Clone(t *testing.T) {
	clone, _ := obj1.Clone()
	t.Log(clone.ToString())
	if obj1 == clone {
		t.Error("obj1 and obj2 should not be equal")
	}
}
