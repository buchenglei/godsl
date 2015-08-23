package basic

import "testing"
import "fmt"

var (
	err        error
	c          int
	s          int
	f          int
	r          int
	data       interface{}
	funcStatus bool
)

func Test_Queue(t *testing.T) {
	q := NewQueue(5)
	if q.IsEmpty() == true && q.IsFull() == false {
		t.Log("  NewQueue success!")
	} else {
		t.Error("  NewQueue error")
	}
	c, s, f, r = q.GetQueueInfo()
	if c == 5 && s == 0 && f == 0 && r == 0 {
		t.Log("  GetQueueInfo success")
	} else {
		t.Error("  GetQueueInfo")
	}

	err = q.Enqueue("test...1")
	_, s, f, r = q.GetQueueInfo()
	if err == nil && s == 1 && f == 0 && r == 1 {
		t.Log("  Enqueue success!")
	} else {
		t.Error("  Enqueue error!")
	}

	err = q.Dequeue()
	_, s, f, r = q.GetQueueInfo()

	if err == nil && s == 0 && f == 1 && r == 1 {
		t.Log("  Dequeue success!")
	} else {
		t.Error("  Dequeue error!")
	}

	// 测试队列满的情况
	err = q.Enqueue("test...1")
	err = q.Enqueue("test...2")
	err = q.Enqueue("test...3")
	err = q.Enqueue("test...4")
	err = q.Enqueue("test...5")
	_, s, f, r = q.GetQueueInfo()
	if s == 5 && f == 1 && r == 1 {
		t.Log("  Full Queue success!")
	} else {
		t.Error("  Full Queue error!")
		fmt.Println(s, f, r)
	}

	// 注意这里只是将队列的第一个值返回，并未从队列中移除
	// 队列并没有发生变化
	data, _ = q.Front()
	if data == "test...1" {
		t.Log("  Front success!")
	} else {
		t.Error("  Front error!")
	}

	data, err = q.FrontAndDequeue()
	_, s, f, r = q.GetQueueInfo()
	if err == nil && data == "test...1" && s == 4 && f == 2 && r == 1 {
		t.Log("  FrontAndDequeue success!")
	} else {
		t.Error("  FrontAndDequeue error!")
	}

	q.EmptyQueue()
	_ = q.Enqueue("test...1")
	_ = q.Enqueue("test...2")
	_ = q.Enqueue("test...3")
	_ = q.Enqueue("test...4")
	_ = q.Enqueue("test...5")
	q.ForceEnqueue("test...6")
	data, _ = q.Front()
	_, s, f, r = q.GetQueueInfo()
	if data == "test...2" && s == 5 && f == 1 && r == 1 {
		t.Log("  ForceEnqueue success!")
	} else {
		t.Error("  ForceEnqueue error!")
	}

	if q.IsFull() {
		t.Log("  The Queue is full!")
	} else {
		t.Error("  The Queue should be full, but not!")
	}

	err = q.Enqueue("test...7")
	if err != nil {
		t.Log("  " + err.Error() + "!")
	} else {
		t.Error("  The Queue should be full, but not!")
	}

	q.EmptyQueue()
	_ = q.Enqueue("test...1")
	_ = q.Enqueue("test...2")
	_ = q.Enqueue("test...3")
	_ = q.Enqueue("test...4")
	_ = q.Enqueue("test...5")
	_ = q.Dequeue()
	_ = q.Dequeue()
	_ = q.Dequeue()
	q.AddCapacity(3)
	c, s, f, r = q.GetQueueInfo()
	data, _ = q.Front()
	if data == "test...4" && c == 8 && s == 2 && f == 6 && r == 0 {
		t.Log("  AddCapacity success!")
	} else {
		t.Error("  AddCapacity error!")
	}

}
