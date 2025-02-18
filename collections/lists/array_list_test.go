package lists

import (
	"fmt"
	"strings"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	list := NewArrayListWithCapacity[string](5)
	list.Add("Java")
	list.Add("Go")
	list.AddAtIndex(1, "C++")

	fmt.Println(list.Size())
	fmt.Println(list.Get(1))
	fmt.Println(list.Contains("Go"))

	list.Remove(0)
	fmt.Println(list.IndexOf("C++"))
}

func TestArrayList_AddAtIndex(t *testing.T) {
	testCases := []struct {
		name     string
		list     *ArrayList[int]
		index    int
		newValue int
	}{
		{
			name:     "add zero index",
			list:     NewArrayList[int](),
			index:    0,
			newValue: 100,
		},
		{
			name:     "add at index",
			list:     NewArrayListOf[int]([]int{1, 2, 3}),
			index:    2,
			newValue: 100,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.list.AddAtIndex(tc.index, tc.newValue)
			for v := range tc.list.Iterator() {
				fmt.Println(v)
			}
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	testCases := []struct {
		name     string
		list     *ArrayList[int]
		index    int
		expected int
	}{
		{
			name:     "get index",
			list:     NewArrayListOf[int]([]int{3, 2, 5, 8, 4, 7, 6, 9}),
			index:    2,
			expected: 5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			get := tc.list.Get(tc.index)
			t.Log(get)
			if get != tc.expected {
				t.Errorf("got %d, want %d", get, tc.expected)
			}
		})
	}
}

func TestArrayList_Set(t *testing.T) {
	list1 := NewArrayListOf[int]([]int{3, 2, 5, 8, 4, 7, 6, 9})
	list1.Set(2, 100)
	get := list1.Get(2)
	t.Log(get)
	if get != 100 {
		t.Errorf("got %d, want %d", get, 100)
	}

	list2 := NewArrayListOf[string]([]string{"name", "age", "sex"})
	list2.Set(2, "tags")
	s := list2.Get(2)
	t.Log(s)
	if strings.Compare(s, "tags") != 0 {
		t.Errorf("got %s, want %s", s, "tags")
	}
}

func TestArrayList_Remove(t *testing.T) {
	arrayList := NewArrayListOf([]int{3, 2, 5, 8, 4, 7, 6, 9})
	arrayList.Remove(3)
	t.Log(arrayList.Size())
	t.Log(arrayList.elements)
}

func TestArrayList_Size(t *testing.T) {
	arrayList := NewArrayListOf[int]([]int{3, 2, 5, 8, 4, 7, 6, 9})
	t.Log(arrayList.Size())
}

func TestArrayList_IsEmpty(t *testing.T) {
	list1 := NewArrayListOf[int]([]int{3, 2, 5, 8, 4, 7, 6, 9})
	t.Log(list1.IsEmpty())
	list2 := NewArrayList[int]()
	t.Log(list2.IsEmpty())
}

func TestArrayList_Contains(t *testing.T) {
	list1 := NewArrayListOf[int]([]int{3, 2, 5, 8, 4, 7, 6, 9})
	t.Log(list1.Contains(9))
	t.Log(list1.Contains(999))
}

func TestArrayList_IndexOf(t *testing.T) {
	list1 := NewArrayListOf[int]([]int{3, 2, 5, 8, 4, 7, 6, 9})
	t.Log(list1.IndexOf(9))
	t.Log(list1.IndexOf(999))
}

func TestArrayList_Clear(t *testing.T) {
	list1 := NewArrayListOf[int]([]int{3, 2, 5, 8, 4, 7, 6, 9})
	t.Log(list1.elements)
	list1.Clear()
	t.Log(list1.elements)
}

func TestArrayList_Iterator(t *testing.T) {
	list1 := NewArrayListOf[int]([]int{3, 2, 5, 8, 4, 7, 6, 9})
	for v := range list1.Iterator() {
		t.Log(v)
	}
}

func TestArrayList_Sort(t *testing.T) {
	list1 := NewArrayListOf[int]([]int{3, 2, 5, 8, 4, 7, 6, 9})
	list1.Sort(func(a, b int) bool {
		return a < b
	})
	t.Log(list1.elements)
}
