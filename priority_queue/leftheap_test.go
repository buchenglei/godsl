package priority_queue

import (
	"fmt"
	"testing"
)

// 需要说明的是：
// 按这里插入的顺序，并不能得到书中例子的树结构
// 进过反复考量，我认为代码并没有出错
func Test_LeftHeap(t *testing.T) {
	H := NewLeftHeap(3)

	H = Insert(10, H)
	H = Insert(8, H)
	H = Insert(21, H)
	H = Insert(14, H)
	H = Insert(17, H)
	H = Insert(23, H)
	H = Insert(26, H)
	if H.left.element == 8 && H.right.element == 10 &&
		H.left.right.element == 14 && H.right.right.left.element == 26 {
		t.Log("  Insert success!")
	} else {
		t.Error("  Insert error! --  ", H.left.element)
	}
	
	H, min := DeleteMin(H)
	if min == 3 && H.element == 8 && H.left.element == 10 && 
		H.right.element == 21 && H.left.right.left.left.element == 26 {
		t.Log("  DeleteMin success!")
	} else {
		t.Error("  DeleteMin error!")
	}
	
	PrintLeftHeap(H)
	println()
}

// 后序遍历
func PrintLeftHeap(H PriorityQueue) {
	if H == nil {
		return
	}

	PrintLeftHeap(H.left)
	PrintLeftHeap(H.right)
	fmt.Print(H.element, "  ")
}
