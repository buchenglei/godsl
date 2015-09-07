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

// 创建一个新的节点
func newNode(data interface{}) *AvlNode {
	node := new(AvlNode)
	node.data = data
	node.left = nil
	node.right = nil
	node.height = 0
	return node
}

// 创建一棵新的AVL树
func NewAvlTree(data interface{}, f comparator) (*AvlNode, error) {
	if data == nil && f == nil {
		return nil, fmt.Errorf("Need data or compare function!")
	}
	compare = f
	return newNode(data), nil

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

	K2.height = max(height(K2.left), height(K2.right)) + 1
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
	K3.right = singleRotateWithLeft(K3.right)

	return singleRotateWithRight(K3)
}

// 将值插入AVL树中
// 因为该方法会自动调整树的结构
// 所以树的根节点有可能会被旋转到其他的地方
// 所以，每次调用该方法是。应该用如下的方法
// Tree = Tree.Insert(2)
func (n *AvlNode) Insert(data interface{}) *AvlNode {
	switch compare(data, n.data) {
	case -1:
		if n.left != nil {
			n.left = n.left.Insert(data)
			// 判断是否满足平衡条件
			if height(n.left)-height(n.right) == 2 {
				// 判断节点插入的位置在左节点的左边还是右边
				if compare(data, n.left.data) == -1 {
					// 在左边，执行左单旋转
					n = singleRotateWithLeft(n)
				} else {
					// 在右边，执行左双旋转
					n = doubleRotateWithLeft(n)
				}
			}
		} else {
			n.left = newNode(data)
		}
	case 1:
		if n.right != nil {
			n.right = n.right.Insert(data)
			if height(n.right)-height(n.left) == 2 {
				if compare(data, n.right.data) == 1 {
					n = singleRotateWithRight(n)
				} else {
					n = doubleRotateWithRight(n)
				}
			}
		} else {
			n.right = newNode(data)
		}
	}
	n.height = max(height(n.left), height(n.right)) + 1
	return n
}

// 返回树中的最小节点
func (n *AvlNode) FindMin() *AvlNode {
	var finded *AvlNode
	if n.left != nil {
		finded = n.left.FindMin()
	} else {
		finded = n
	}

	return finded
}

// 返回树中的最大的节点
func (n *AvlNode) FindMax() *AvlNode {
	var finded *AvlNode
	if n.right != nil {
		finded = n.right.FindMax()
	} else {
		finded = n
	}

	return finded
}

// 查找指定的节点
func (n *AvlNode) Find(data interface{}) *AvlNode {
	var finded *AvlNode = nil
	switch compare(data, n.data) {
	case -1:
		if n.left != nil {
			finded = n.left.Find(data)
		}
	case 1:
		if n.right != nil {
			finded = n.right.Find(data)
		}
	case 0:
		return n
	}

	return finded

}

// 返回当前节点的值
func (n *AvlNode) GetData() interface{} { return n.data }

// 重新设置当前节点的值
func (n *AvlNode) SetData(data interface{}) { n.data = data }

// 返回当前节点的左节点
func (n *AvlNode) Left() *AvlNode { return n.left }

// 返回当前节点的右节点
func (n *AvlNode) Right() *AvlNode { return n.right }
