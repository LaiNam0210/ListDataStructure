package list

import (
	"errors"
	"reflect"
	"runtime"
)

var (
	defaultSizeQueue = 4096
)

type Queue interface {
	// basic operations:
	EnQueue(interface{}) error
	EnQueues(...interface{}) error
	DeQueue() interface{}
	DeQueues(int) []interface{}
	QueueFront() interface{}
	QueueRear() interface{}

	// extend operations:
	IsFull() bool
	IsEmpty() bool
	Size() int64
	Clear()
	Total() int64
	Count() int64
}

type QueuePtr struct {
	total int64
	size  int64
	count int64

	front *node
	rear  *node
}

func (q *QueuePtr) Count() int64 {
	return q.count
}
func CreateQueuePtr(args ...interface{}) Queue {
	queue := &QueuePtr{}

	if len(args) == 0 {
		queue.size = int64(defaultSizeQueue)
		return queue
	}

	if len(args) == 1 {
		switch r := args[0].(type) {
		case int:
			queue.size = int64(r)
			return queue
		}
	}

	return nil
}

func (q *QueuePtr) EnQueue(element interface{}) (err error) {
	if q.IsFull() {
		return errors.New("Queue is Full")
	}

	node := new(node)
	node.data = element

	if q.rear == nil {
		q.rear = node
		q.front = node
	} else {
		q.rear.next = node
		q.rear = node
	}

	q.count++
	q.total++
	return nil
}

func (q *QueuePtr) EnQueues(elements ...interface{}) (err error) {

	for i := 0; i < len(elements); i++ {
		Val := reflect.ValueOf(elements[i])

		switch Val.Kind() {

		case reflect.Slice, reflect.Array:

			for a := 0; a < Val.Len(); a++ {
				err = q.EnQueue(Val.Index(a).Interface())
				if err != nil {
					i = len(elements)
					break
				}
			}
		case reflect.Map:

			mapKeys := Val.MapKeys()
			for i := 0; i < len(mapKeys); i++ {

				mapKey := mapKeys[i]

				err = q.EnQueue(mapKey.Interface())
				if err != nil {
					i = len(elements)
				}
				err = q.EnQueue(Val.MapIndex(mapKey).Interface())
				if err != nil {
					break
					i = len(elements)

				}
			}

		default:
			err = q.EnQueue(elements[i])
			if err != nil {
				i = len(elements)
			}

		}
	}
	return
}
func (q *QueuePtr) DeQueue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	Node := q.front
	q.front = Node.next
	if q.count == 1 {
		q.rear = q.front
	}
	q.count--
	return Node.data

}

func (q *QueuePtr) DeQueues(number int) (elements []interface{}) {
	if number < 0 {
		return
	}
	for i := 0; i < number; i++ {

		r := q.DeQueue()
		if r == nil {
			break
		}
		elements = append(elements, r)
	}
	return

}
func (q *QueuePtr) QueueFront() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.front.data
}

func (q *QueuePtr) QueueRear() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.rear.data

}

func (q *QueuePtr) Total() int64 {
	return q.total
}

func (q *QueuePtr) Size() int64 {
	return q.size
}
func (q *QueuePtr) Clear() {
	q.rear = nil
	q.front = nil
	q.count = 0
	runtime.GC()

}

func (q *QueuePtr) IsEmpty() bool {
	if q.count == 0 {
		return true
	}
	return false
}

func (q *QueuePtr) IsFull() bool {
	if q.count == q.size {
		return true
	}
	return false
}
