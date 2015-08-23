package basic

import "fmt"

type Queue struct {
	capacity int // 队列的容量
	size     int // 队列中实际数据的个数
	front    int // 指向队首的指针
	rear     int // 指向队尾最后一个元素的下一个空位的指针,
	array    []interface{}
}

// 创建一个新的队列
func NewQueue(capacity int) *Queue {
	q := new(Queue)
	q.capacity = capacity
	q.size = 0
	q.front = 0
	q.rear = 0
	q.array = make([]interface{}, capacity)
	return q
}

// 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// 判断队列是否已满
func (q *Queue) IsFull() bool {
	return q.size == q.capacity
}

// 返回队列的相关信息
// c - capacity
// s - size
// f - front
// r - rear
func (q *Queue) GetQueueInfo() (c int, s int, f int, r int) {
	c = q.capacity
	s = q.size
	f = q.front
	r = q.rear
	return
}

// 将数据插入至队尾
func (q *Queue) Enqueue(data interface{}) error {
	if q.IsFull() {
		return fmt.Errorf("The Queue is full")
	}

	// rear指向的内存单元始终是可写入的
	q.array[q.rear] = data
	q.size++
	// 如果当前的rear指向数组的尾部,直接 +1 的话
	// 将会导致数组下标越界，则需要将rear置为0
	// 这是在队列未满的情况下才执行的过程
	if q.rear == q.capacity-1 {
		q.rear = 0
	} else {
		q.rear++
	}
	return nil
}

// 丢弃队首的数据,并将队首的指针指向下一个数据
func (q *Queue) Dequeue() error {
	if q.IsEmpty() {
		return fmt.Errorf("The Queue is empty")
	}

	q.size--
	// 如果当前的front指向数组的尾部,直接 +1 的话
	// 将会导致数组下标越界，则需要将front置为0
	// 这是在队列不为空的情况下才执行的过程
	if q.front == q.capacity-1 {
		q.front = 0
	} else {
		q.front++
	}
	return nil

}

// 返回队首数据,但是front指针并不向后移动
// 仅用于查看队首的数据
func (q *Queue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("The Queue is empty")
	}

	return q.array[q.front], nil
}

// 返回队首数据并将front指针后移一位
func (q *Queue) FrontAndDequeue() (interface{}, error) {
	data, err := q.Front()
	_ = q.Dequeue()
	return data, err
}

// 将数据强制插入队尾
// 执行该函数时无需考虑队列满的情况
// 因为当队列满的情况下执行该函数
// 该函数会自动将队首数据丢弃
func (q *Queue) ForceEnqueue(data interface{}) {
	if q.IsFull() {
		q.Dequeue()
	}

	q.Enqueue(data)
}

// 将该队列重新置为空
func (q *Queue) EmptyQueue() {
	q.size = 0
	q.front = 0
	q.rear = 0
}

// 将原队列的容量扩大指定的值
func (q *Queue) AddCapacity(num int) error {
	// 队列的容量只能扩大，不能缩小
	if num <= 0 {
		return fmt.Errorf("The Added value must be more than 0")
	}

	newArray := make([]interface{}, q.capacity+num)
	copy(newArray, q.array)
	q.array = newArray
	if q.rear < q.front {
		// 如果rear在front的前面
		// 则要将从front到原数组的尾部的值整体向后移动num个单位
		// i就是原数组最大的index
		for i := q.capacity - 1; i >= q.front; i-- {
			// 将满足条件的每一个数据向后移动num、个单位
			q.array[i+num] = q.array[i]
		}
		// 不能忘记front指针也要向后移动
		q.front += num
	}
	// 保存当前数组的实际容量
	q.capacity += num
	return nil
}
