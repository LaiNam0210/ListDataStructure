package list

import (
	"testing"
)

//testcase1
func TestQueueArray(T *testing.T) {
	queue := CreateQueuePtr(1000)

	arr := []interface{}{"One", 1, "Two", 2, "Three", 3, "Four", 4, "Five", 5}

	for i := 0; i < len(arr); i++ {
		queue.EnQueue(arr[i])
	}

	expected := []interface{}{5, "Five", 4, "Four", 3, "Three", 2, "Two", 1, "One"}

	for i := len(arr); i > 0; i-- {
		if e := queue.DeQueue(); e != expected[i-1] {
			T.Errorf("Expected %v But Got %v", expected[i], e)
		}
	}
	if r := queue.Count(); r != 0 {
		T.Errorf("Expected %v But Got %v", 0, r)
	}
}

//testcase Array,map
func TestQueueInsertArrayMap(T *testing.T) {
	queue := CreateQueuePtr(1000)
	arr := []interface{}{"One", 1, "Two", 2, "Three", 3, "Four", 4, "Five", 5}

	queue.EnQueues(arr)

	for i := 0; i < len(arr); i++ {
		if r := queue.DeQueue(); r != arr[i] {
			T.Errorf("Expected %v But Got %v", arr[i], r)
		}
	}
	if r := queue.Count(); r != 0 {
		T.Errorf("Expected %v But Got %v", 0, r)
	}

	values := make(map[interface{}]interface{})

	for i := 0; i < len(arr); i += 2 {
		values[i] = arr[i]
	}
	/*
		Value[0] = One
		Value[1] = Two
		Value[2] = Three
		Value[3] = Four
		Value[4] = Five
	*/
	queue.EnQueues(values)
	for i := 0; i < len(arr); i += 2 {
		if e := queue.DeQueues(2); e[0] != arr[i+1] && e[1] != arr[i] {
			T.Errorf("Expected %v,%v But Got %v,%v", arr[i+1], arr[i], e[0], e[1])
		}
	}

}

//testcase 3
func TestCase3(T *testing.T) {
	queue := CreateQueuePtr()

	arr := []int{1, 2, 3, 4, 5}
	queue.EnQueues(arr)

	r := queue.DeQueues(100)

	for i := 0; i < len(arr); i++ {
		if arr[i] != r[i] {
			T.Errorf("Expected %v But Got %v", arr[i], r[i])
		}
	}
}

//testCase4
func TestCase4(T *testing.T) {
	queue := CreateQueuePtr(12)
	arr := make(map[int]int)
	arr[1] = 100
	arr[2] = 200
	arr[3] = 300
	arr[4] = 400
	arr[5] = 500
	arr[6] = 600

	queue.EnQueues(arr)

	r := queue.DeQueues(12)
	/*
		r: slice {1,100,2,200,3,300,4,400,5,500,6,600}
	*/
	for i := 0; i < 12; i += 2 {
		k := i/2 + 1
		if r[i] != k && r[i+1] != arr[k] {
			T.Errorf("Expected %v,%v But Got %v,%v", k, arr[k], r[i], r[i+1])
		}
	}

}
