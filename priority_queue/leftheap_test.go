package priority_queue

import (
	"fmt"
	"testing"
)

func Test_LeftHeap(t *testing.T) {
	H := NewLeftHeap(3)

	H = Insert(10, H)
	H = Insert(8, H)
	H = Insert(21, H)
	H = Insert(14, H)
	H = Insert(17, H)
	H = Insert(23, H)
	H = Insert(26, H)
	/*if H.left.element == 10 {
		t.Log("  Insert success!")
	} else {
		t.Error("  Insert error! --  ", H.left.element)
	}*/
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
