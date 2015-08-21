package basic

import (
	"testing"
)

const (
	STACK_LENGTH = 5
)

var (
	err        error
	data       interface{}
	funcStatus bool
)

func Test_Stack(t *testing.T) {
	s := NewStack(STACK_LENGTH)
	c, top, fs := s.GetStackInfo()
	if c == 5 && top == -1 && fs == 5 {
		t.Log("  CreateNewStack and GetStackInfo success!")
	} else {
		t.Error("  CreateNewStack or GetStackInfo error!")
	}

	isempty := s.IsEmpty()
	if isempty {
		t.Log("  IsEmpty success!")
	} else {
		t.Error("  IsEmpty error!")
	}

	isfull := s.IsFull()
	if isfull == false {
		t.Log("  IsFull success!")
	} else {
		t.Error("  IsFull error!")
	}

	err = s.Push("test...1")
	_, top, fs = s.GetStackInfo()
	if err == nil && top == 0 && fs == 4 {
		t.Log("  Push success!")
	} else {
		t.Error("  Push error!")
	}

	err = s.Pop()
	_, top, fs = s.GetStackInfo()
	if err == nil && top == -1 && fs == 5 {
		t.Log("  Pop success!")
	} else {
		t.Error("  Pop error!")
	}

	_ = s.Push("test...2")
	data, err = s.TopAndPop()
	_, top, fs = s.GetStackInfo()
	if data == "test...2" && err == nil && top == -1 && fs == 5 {
		t.Log("  TopAndPop success")
	} else {
		t.Error("  TopAndPop error!")
	}

	_ = s.Push("test...1")
	_ = s.Push("test...2")
	_ = s.Push("test...3")
	err = s.PopAll()
	_, top, fs = s.GetStackInfo()
	if err == nil && top == -1 && fs == 5 {
		t.Log("  PopAll success!")
	} else {
		t.Error("  PopAll error!")
	}

	// 此时栈为空
	_ = s.Push("test...1")
	_ = s.Push("test...2")
	_ = s.Push("test...3")
	// 增加栈空间
	err = s.AddCapacity(4)
	c, top, fs = s.GetStackInfo()
	if err == nil && c == 9 && top == 2 && fs == 6 {
		funcStatus = true
	} else {
		funcStatus = false
	}

	err = s.AddCapacity(-7)
	c, top, fs = s.GetStackInfo()
	if err == nil && c == 2 && top == 1 && fs == 0 {
		funcStatus = true
	} else {
		funcStatus = false
	}

	if funcStatus {
		t.Log("  AddCapacity success!")
	} else {
		t.Error("  AddCapacity error!")
	}
}
