// 使用数组实现的栈结构
// 由于存储的存储使用的是数组，所以并不依赖前面的链表
package basic

import (
	"fmt"
)

// 声明栈的结构
type Stack struct {
	capacity int           // 栈的容量
	top      int           // 指向栈顶的指针,index从0开始
	array    []interface{} // 用于保存数据的数组
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

// 由于top指针的index从0开始
// 所以当栈满时top的值就等于capacity - 1
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

// 弹出栈顶的值并返回
func (s *Stack) TopAndPop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("The stack is empty!")
	}

	data := s.array[s.top]
	s.top--
	return data, nil
}

// 将栈置为空， top = -1
func (s *Stack) PopAll() error {
	if s.IsEmpty() {
		return fmt.Errorf("The stack is empty!")
	}

	s.top = -1
	return nil
}

// 增加当前栈的容量
// 5(正数) 将栈的容量设为: 原来的值 + 5
// -5(负数) 将站的容量设为: 原来的值 - 5
// 有一点需要注意的是：
// 减少当前栈的容量，若新栈的容量不足以保存之前栈的所有值
// 则丢弃原栈多出去的值丢弃，并将top指向新栈的栈顶
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
