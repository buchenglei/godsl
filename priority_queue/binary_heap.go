// 二叉堆(binary heap)的结构实现
// 根据堆序性质，最小元总是可以在根处找到
package priority_queue

import "fmt"

// 定义一个用于比较二叉堆中元素大小的函数
// 0	相等
// 1	a > b
// -1	a < b
// -2	error
type CompareFunc func(a interface{}, b interface{}) int

// struct of binary heap
type BinHeap struct {
	// capacity of whole heap
	capacity int
	// current size
	size int
	// a slice to store element
	elements []interface{}
	// compare function
	compare CompareFunc
}

// 创建一个新的二叉堆
// unreachables是一个不可达的最小值
// comp 用于比较二叉堆中两个元素的大小
func NewBinHeap(c int, unreachable interface{}, comp CompareFunc) (*BinHeap, error) {
	if c < 3 {
		return nil, fmt.Errorf("Heap size is too small")
	}

	heap := new(BinHeap)
	heap.capacity = c
	heap.size = 0
	heap.compare = comp
	// actual capacity of slice equals c + 1
	// beacuse slice[0] = unreachable
	heap.elements = make([]interface{}, c+1)
	heap.elements[0] = unreachable

	return heap, nil
}

func (b *BinHeap) IsFull() bool { return b.size == b.capacity }

func (b *BinHeap) IsEmpty() bool { return b.size == 0 }

func (b *BinHeap) Insert(element interface{}) (bool, error) {
	if b.IsFull() {
		return false, fmt.Errorf("Binary Heap is full")
	}

	b.size += 1
	var i int // 循环外还需要使用
	// 始终是 i 的父节点和待插入的元素比较大小
	for i = b.size; b.compare(b.elements[i/2], element) == 1; /* > */ i /= 2 {
		b.elements[i] = b.elements[i/2]
	}
	// 当循环结束后，slice中的的空出来的位置就是element该插入的地方
	// i 始终指向slice中空出来的那个位置
	b.elements[i] = element

	return true, nil
}

// 删除二叉堆中最小的元素（即 根几点）
func (b *BinHeap) DeleteMin() interface{} {
	if b.IsEmpty() {
		return nil
	}

	// 获得最小值，在树的根处
	minElement := b.elements[1]
	// 删除最小元素后，将最后一个元素作为游离的元素
	// 插入到堆中的合适的位置
	lastElement := b.elements[b.size]
	b.size -= 1

	var child, i int
	for i = 1; i*2 <= b.size; i = child {
		// 1.find a smaller child
		child = i * 2 // left child
		// 2.判断右节点是否比左节点小
		if child != b.size &&
			b.compare(b.elements[child+1], b.elements[child]) == -1 /* < */ {
			child += 1
		}

		// 3.将child所指向的元素和游离出的last元素比较
		// 如果lastElement大于child所指向的元素
		// 那么将child所指向的元素赋值给 i 指向的位置
		// 并将child的指向下移一层继续比较
		// 直到lastElement比child所指向的元素小
		// 那么将lastElement插入到 i 所指向的位置
		if b.compare(lastElement, b.elements[child]) == 1 /* > */ {
			b.elements[i] = b.elements[child]
		} else {
			break
		}
	}
	// 当循环结束的时候 i 所指向的位置就是放置lastElementd额位置
	b.elements[i] = lastElement

	return minElement
}

// 获得二叉堆中的最小元素，但不删除它
func (b *BinHeap) FindMin() interface{} {
	if b.IsEmpty() {
		return nil
	} else {
		return b.elements[1]
	}
}

// 清空当前的二叉堆
func (b *BinHeap) Empty() {
	for i := 1; i <= b.size; i++ {
		b.elements[i] = nil
	}

	b.size = 0
}
