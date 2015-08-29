package tree

import (
	"fmt"
	"testing"
)

func Test_SearchTree(t *testing.T) {
	// 创建一个新的查找树
	// 并设置比较函数
	st := NewSearchTree(10, _compare)
	st.Insert(5)
	st.Insert(8)
	st.Insert(11)
	st.Insert(12)
	min := st.FindMin().GetData()
	max := st.FindMax().GetData()
	if min == 5 {
		t.Log("  FindMin success!")
	} else {
		t.Error("  FindMin error!")
	}

	if max == 12 {
		t.Log("  FindMax success!")
	} else {
		fmt.Println("The max is ", max)
		t.Error("  FindMax error!")
	}

	finded := st.Find(8).GetData()
	if finded == 8 {
		t.Log("  Find success!")
	} else {
		t.Error("  find error!")
	}

	st.Delete(5)
	st.Delete(12)
	min = st.FindMin().GetData()
	max = st.FindMax().GetData()
	if min == 8 && max == 11 {
		t.Log("  Delete success!")
	} else {
		t.Error("  Delete error!")
	}

	// 遍历输出树
	PrintTree(st)

}

func PrintTree(st *SearchTree) {
	if st.left != nil {
		PrintTree(st.left)
	}

	if st.right != nil {
		PrintTree(st.right)
	}

	fmt.Println(st.data)
}

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
