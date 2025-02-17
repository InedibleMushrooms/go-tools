package bases

import "testing"

type Person struct {
	BaseObject
	Name string
	Tags []string
}

func TestBaseObject_Equals(t *testing.T) {
	obj1 := &Person{
		Name: "Alice",
		Tags: []string{"tag1", "tag2"},
	}
	obj2 := &Person{
		Name: "Bob",
		Tags: []string{"tag3", "tag4"},
	}
	if obj1.Equals(obj2) {
		t.Error("obj1 and obj2 should not be equal")
	}
}
