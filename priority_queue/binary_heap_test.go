// 二叉堆的测试文件
package priority_queue

import (
	"fmt"
	"testing"
)

var TestData = [11]int{13, 14, 16, 24, 21, 19, 68, 65, 26, 32, 31}

func Test_BinaryHeap(t *testing.T) {
	heap, err := NewBinHeap(15, -1, _compare)
	if err == nil && heap != nil {
		t.Log("  NewBinHeap success!")
	} else {
		t.Error("  NewBinHeap error!")
	}

	if heap.IsFull() == false && heap.IsEmpty() == true {
		t.Log("  IsFull and IsEmpty success!")
	} else {
		t.Error("  IsFull or IsEmpty error!!")
	}

	heap.Insert(13)
	heap.Insert(21)
	heap.Insert(16)
	heap.Insert(24)
	heap.Insert(31)
	heap.Insert(19)
	heap.Insert(68)
	heap.Insert(65)
	heap.Insert(26)
	heap.Insert(32)
	heap.Insert(14)
	if check_sort(heap) {
		t.Log("  Insert success!")
	} else {
		t.Error("  Insert error!")
	}

	if heap.DeleteMin() == 13 && heap.DeleteMin() == 14 &&
		heap.DeleteMin() == 16 && heap.DeleteMin() == 19 {
		t.Log("  DeleteMin success!")
	} else {
		t.Error("  DeleteMin error!")
	}

	if heap.FindMin() == 21 {
		t.Log("  FindMin success!")
	} else {
		t.Error("  FindMin error!")
		fmt.Println("FindMin is ", heap.FindMin())
	}

	heap.Empty()
	if heap.elements[1] == nil && heap.size == 0 {
		t.Log("  Empty success!")
	} else {
		t.Error("  Empty error!")
	}
}

// 用于比较数据大小的测试函数
// 这里使用整数测试
func _compare(a, b interface{}) int {
	var newA, newB int
	var ok bool
	if newA, ok = a.(int); !ok {
		return -2
	}
	if newB, ok = b.(int); !ok {
		return -2
	}
	if newA > newB {
		return 1
	}
	if newA == newB {
		return 0
	}
	if newA < newB {
		return -1
	}
	// impossible
	return -2
}

func check_sort(b *BinHeap) bool {
	var err bool = true
	for i := 1; i <= b.size; i++ {
		if v, ok := b.elements[i].(int); ok && v != TestData[i-1] {
			fmt.Printf("Error:At %d, Heap is %d, but i want is %d\n", i, v, TestData[i-1])
			err = false
		}
	}

	return err
}

/*
 写在最后
 在测试的过程中一个小问题：
 就是在for循环外 var i int
 但是在for循环中顺手写成for i := 1 .....
 稍不注意就改变了 i 的作用域，本来for循环中的 i
 是要拿到外面用的，现在得到的 i 始终是 0
 遇到的两个小bug都是由于这个原因造成的
 一定要细心 ^_^
*/
