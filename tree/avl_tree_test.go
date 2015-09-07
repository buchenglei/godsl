package tree

import (
	"fmt"
	"testing"
)

func Test_AvlNode(t *testing.T) {
	Tree, _ := NewAvlTree(3, _compare)
	Tree = Tree.Insert(2)
	Tree = Tree.Insert(1)
	Tree = Tree.Insert(4)
	Tree = Tree.Insert(5)
	Tree = Tree.Insert(6)
	Tree = Tree.Insert(7)
	Tree = Tree.Insert(16)
	Tree = Tree.Insert(15)
	Tree = Tree.Insert(14)
	Tree = Tree.Insert(13)
	Tree = Tree.Insert(12)
	Tree = Tree.Insert(11)
	Tree = Tree.Insert(10)
	Tree = Tree.Insert(9)
	Tree = Tree.Insert(8)

	// 遍历输出AVL树
	fmt.Println("Correct order:")
	fmt.Println("1  3  2  5  6  4  8  10  9  12  11  14  16  15  13  7")
	fmt.Println("Actual order:")
	print(Tree)
	fmt.Println()
	fmt.Println("只有在上面的顺序是正确的情况下，后面的测试才有意义")

	min := Tree.FindMin().GetData()
	if min == 1 {
		t.Log("  FindMin success!")
	} else {
		t.Error("  FindMin error!")
	}

	max := Tree.FindMax().GetData()
	if max == 16 {
		t.Log("  FindMax success!")
	} else {
		t.Error("  FindMax error!")
	}

	var randData [3]interface{}
	randData[0] = Tree.Find(10).GetData()
	randData[1] = Tree.Find(3).GetData()
	randData[2] = Tree.Find(7).GetData()
	if randData[0] == 10 && randData[1] == 3 && randData[2] == 7 {
		t.Log("  Find success!")
	} else {
		t.Error("  Find error!")
	}

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

func print(t *AvlNode) {
	if t.left != nil {
		print(t.left)
	}

	if t.right != nil {
		print(t.right)
	}

	fmt.Printf("%d  ", t.GetData())
}
