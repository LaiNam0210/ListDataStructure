package main

import (
	"algorithm/list"
	"fmt"
)

func main() {
	queue := list.CreateQueuePtr()
	arr := []int{1, 2, 3, 4, 5}

	queue.EnQueues(arr)

	r := queue.DeQueues(5)
	fmt.Println(r)
}
