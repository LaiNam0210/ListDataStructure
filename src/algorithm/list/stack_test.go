package list

import (
	"testing"
)

//testcase1
func TestStackArray(T *testing.T) {
	stack := CreateStackArr(100)

	arr := []interface{}{"One", 1, "Two", 2, "Three", 3, "Four", 4, "Five", 5}

	for i := 0; i < len(arr); i++ {
		stack.Push(arr[i])
	}

	expected := []interface{}{5, "Five", 4, "Four", 3, "Three", 2, "Two", 1, "One"}

	for i := 0; i < len(arr); i++ {
		if e := stack.Pop(); e != expected[i] {
			T.Errorf("Expected %v But Got %v", expected[i], e)
		}
	}

}

func TestStackPointer(T *testing.T) {
	stack := CreateStackPtr()

	arr := []interface{}{"One", 1, "Two", 2, "Three", 3, "Four", 4, "Five", 5}

	for i := 0; i < len(arr); i++ {
		stack.Push(arr[i])
	}

	expected := []interface{}{5, "Five", 4, "Four", 3, "Three", 2, "Two", 1, "One"}

	for i := 0; i < len(arr); i++ {
		if e := stack.Pop(); e != expected[i] {
			T.Errorf("Expected %v But Got %v", expected[i], e)
		}
	}
}
