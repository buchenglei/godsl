// 二叉查找树
// 基本的二叉查找树的实现
package tree

//import "fmt"

type SearchTree struct {
	data  interface{}
	left  *SearchTree
	right *SearchTree
}

// 用于比较节点元素大小的函数类型
// 1	a大于b
// 0	a等于b
// -1	a小于b
// -2	出错
type comparator func(a, b interface{}) int

var (
	compare comparator
)

// 创建一个新的查找树的根节点
// data	根节点的值
// f	该查找树的比较函数
func NewSearchTree(data interface{}, f comparator) *SearchTree {
	compare = f
	return newNode(data)
}

// 创建一个新的节点
func newNode(data interface{}) *SearchTree {
	node := new(SearchTree)
	node.data = data
	node.left = nil
	node.right = nil
	return node
}

// 返回相应值得节点
func (st *SearchTree) Find(data interface{}) *SearchTree {
	tmp := st
	for tmp != nil {
		switch compare(data, tmp.data) {
		case 1:
			tmp = tmp.right
		case 0:
			return tmp
		case -1:
			tmp = tmp.left

		}
	}
	return nil
}

// 返回二叉查找树中最小的节点
func (st *SearchTree) FindMin() *SearchTree {
	for st.left != nil {
		st = st.left
	}
	// 最左边的叶节点
	return st
}

// 返回二叉查找树中最大的节点
func (st *SearchTree) FindMax() *SearchTree {
	for st.right != nil {
		st = st.right
	}
	// 最右边的叶节点
	return st
}

// 按照二叉查找树的规则向树中插入新的节点
func (st *SearchTree) Insert(data interface{}) {
	// 新的节点始终插在叶节点上，并不会在上下两个节点之间进行插入
	switch compare(data, st.data) {
	case 1:
		if st.right != nil {
			st = st.right
			st.Insert(data)
		} else {
			st.right = newNode(data)
		}
	case -1:
		if st.left != nil {
			st = st.left
			st.Insert(data)
		} else {
			st.left = newNode(data)
		}
	}
}

// 删除指定节点
func (st *SearchTree) Delete(data interface{}) {
	switch compare(data, st.data) {
	case 1:
		if st.right != nil {
			st = st.right
			st.Delete(data)
		}
	case -1:
		if st.left != nil {
			st = st.left
			st.Delete(data)
		}
	case 0:
		if st.left != nil && st.right != nil {
			// 匹配到的节点有两个子节点的情况
			tmp := st.FindMin()
			st.data = tmp.data
			st = st.right
			st.Delete(tmp.data)
		} else {
			// 匹配到的节点有一个或零个子节点
			if st.left == nil && st.right != nil {
				st.data = st.right.data
				st.right = nil
			} else if st.right == nil && st.left != nil {
				st.data = st.left.data
				st.left = nil
			} else {
				st.data = nil
			}
		}
	}
	// 每一层递归结束后，当前节点都要处理一下无效节点
	// 无效节点：
	// 就是st.left or st. right != nil，但是其data字段为nil
	if st.left != nil && st.left.data == nil {
		st.left = nil
	}

	if st.right != nil && st.right.data == nil {
		st.right = nil
	}
}

// 当前节点的基本操作
func (st *SearchTree) GetData() interface{} { return st.data }

func (st *SearchTree) GetLeft() *SearchTree { return st.left }

func (st *SearchTree) GetRight() *SearchTree { return st.right }
