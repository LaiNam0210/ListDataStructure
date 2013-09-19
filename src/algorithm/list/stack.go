package list

import (
	"errors"
)

type StackArr struct {
	elements []interface{}
	count    int64
	size     int64
	total    int64
}

var (
	defaultSizeStack = 4096
)

func CreateStackArr(arg ...interface{}) Stack {
	stack := &StackArr{}
	switch len(arg) {
	case 0:
		stack.elements = make([]interface{}, defaultSizeStack)[:0]
		stack.size = int64(defaultSizeStack)
		return stack
	case 1:
		switch r := arg[0].(type) {
		case int:
			if r > 0 {
				stack.elements = make([]interface{}, r)[:0]
				stack.size = int64(r)
				return stack
			}
		}

	default:
		return nil
	}
	return nil
}

func (s *StackArr) Push(element interface{}) (err error) {
	if s.IsFull() {
		err = errors.New("Stack is full")
	}
	s.elements = append(s.elements, element)
	s.count++
	s.total++
	return nil

}

func (s *StackArr) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	s.count--
	return s.elements[s.count]

}

func (s *StackArr) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.elements[s.count]
}

func (s *StackArr) IsEmpty() bool {
	if s.count == 0 {
		return true
	}
	return false
}

func (s *StackArr) IsFull() bool {
	if s.count == s.size {
		return true
	}
	return false
}

func (s *StackArr) Size() int64 {
	return s.size
}

func (s *StackArr) Total() int64 {
	return s.total
}

func (s *StackArr) Count() int64 {
	return s.count
}

type node struct {
	data interface{}
	next *node
}

type StackPtr struct {
	count     int64
	total     int64
	size      int64
	frontNode *node
}

func (s *StackPtr) Push(element interface{}) (err error) {
	if s.IsFull() {
		err = errors.New("stack is Full")
	}
	node := new(node)
	node.data = element

	node.next = s.frontNode
	s.frontNode = node
	s.count++
	s.total++
	return nil
}

func (s *StackPtr) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	Node := s.frontNode

	s.frontNode = Node.next
	s.count--

	return Node.data
}

func (s *StackPtr) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	Node := new(node)
	Node.data = s.frontNode.data
	return Node
}

func (s *StackPtr) IsEmpty() bool {
	if s.count == 0 {
		return true
	}
	return false
}
func (s *StackPtr) IsFull() bool {
	if s.size == s.count {
		return true
	}
	return false
}

func (s *StackPtr) Total() int64 {
	return s.total
}

func (s *StackPtr) Size() int64 {
	return s.size
}
func CreateStackPtr(arg ...interface{}) Stack {
	stack := &StackPtr{}

	switch len(arg) {
	case 0:
		stack.size = int64(defaultSizeStack)
		return stack
	case 1:
		switch r := arg[0].(type) {
		case int:
			stack.size = int64(r)
			return stack
		}
	}
	return nil
}
func (s *StackPtr) Count() int64 {
	return s.count
}

type Stack interface {
	Push(interface{}) error
	Pop() interface{}
	Top() interface{}

	IsFull() bool
	IsEmpty() bool
	Size() int64
	Total() int64
	Count() int64
}
