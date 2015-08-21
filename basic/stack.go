// 使用数组实现的栈结构
// 由于存储的存储使用的是数组，所以并不依赖前面的链表
package basic

import (
	"fmt"
)

type Stack struct {
	capacity int
	top      int
	array    []interface{}
}

func NewStack(capacity int) Stack {
	var s Stack
	s.capacity = capacity
	s.top = -1
	s.array = make([]interface{}, capacity)
	return s
}

func (s *Stack) GetStackInfo() (c int, t int, fs int) {
	// c - Capacity
	// t - Top Of Stack
	// fs - Free Space
	c = s.capacity
	t = s.top
	fs = c - t - 1
	return
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) IsFull() bool {
	return (s.top + 1) == s.capacity
}

func (s *Stack) Push(data interface{}) error {
	if s.IsFull() {
		return fmt.Errorf("The stack is full")
	}

	s.top++
	s.array[s.top] = data
	return nil
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return fmt.Errorf("The stack is empty!")
	}

	s.top--
	return nil
}

func (s *Stack) TopAndPop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("The stack is empty!")
	}

	data := s.array[s.top]
	s.top--
	return data, nil
}

func (s *Stack) PopAll() error {
	if s.IsEmpty() {
		return fmt.Errorf("The stack is empty!")
	}

	s.top = -1
	return nil
}

func (s *Stack) AddCapacity(num int) error {
	if s.capacity+num <= 0 {
		return fmt.Errorf("this value will make stack's capacity less than 0,It would be meaningless")
	}

	s.capacity += num
	newArray := make([]interface{}, s.capacity)
	copy(newArray, s.array)
	if (s.top + 1) > s.capacity {
		s.top = s.capacity - 1
	}
	s.array = newArray
	return nil
}
