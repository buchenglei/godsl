package tree

import "fmt"

type AvlNode struct {
	data   interface{}
	left   *AvlNode
	right  *AvlNode
	height int
}

// 用于比较节点元素大小的函数类型
// 1	a大于b
// 0	a等于b
// -1	a小于b
// -2	出错
type comparator func(a, b interface{}) int

var compare comparator

// 计算AVL节点的高度函数
func height(n *AvlNode) int {
	if n == nil {
		return -1
	} else {
		return n.height
	}
}

func NewAvlTree(data interface{}, f comparator) (*AvlNode, error) {
	if data == nil || f == nil {
		return nil, fmt.Errorf("Need data or compare function!")
	}
	node := new(AvlNode)
	node.data = data
	node.left = nil
	node.right = nil
	node.height = 0
	compare = f
	return node, nil

}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}

}

// 左旋转
func singleRotateWithLeft(K2 *AvlNode) *AvlNode {
	var K1 *AvlNode

	K1 = K2.left
	K2.left = K1.right
	K1.right = K2

	K2.height = max(height(K2.left), height(K2.right)) + 1
	K1.height = max(height(K1.left), K2.height) + 1

	return K1
}

// 右旋转
func singleRotateWithRight(K2 *AvlNode) *AvlNode {
	var K1 *AvlNode

	K1 = K2.right
	K2.right = K1.left
	K1.left = K2

	K2.height = max(height(K1.left), height(K1.right)) + 1
	K1.height = max(height(K1.right), K2.height) + 1

	return K1
}

// 左双旋转
// 双旋转可以拆分成两次单旋转
func doubleRotateWithLeft(K3 *AvlNode) *AvlNode {
	K3.left = singleRotateWithRight(K3.left)

	return singleRotateWithLeft(K3)
}

// 右双旋转
func doubleRotateWithRight(K3 *AvlNode) *AvlNode {
	K3.right = singleRotateWithRight(K3.right)

	return singleRotateWithRight(K3)
}
