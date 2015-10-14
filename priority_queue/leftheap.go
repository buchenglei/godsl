package priority_queue

type ElementType int

type TreeNode struct {
	element ElementType
	left    *TreeNode
	right   *TreeNode
	npl     int
}

type PriorityQueue *TreeNode

func NewLeftHeap(element ElementType) PriorityQueue {
	H := new(TreeNode)

	H.element = element
	H.left = nil
	H.right = nil
	H.npl = 0

	return PriorityQueue(H)
}

// 合并左式堆的驱动程序
func Merge(H1, H2 PriorityQueue) PriorityQueue {
	if H1 == nil {
		return H2
	}
	if H2 == nil {
		return H1
	}
	// 左式堆也是满足堆序性质的
	// 所以始终是根元素大的合并到根元素小的堆中
	if H1.element < H2.element {
		return Merge1(H1, H2)
	} else {
		return Merge1(H2, H1)
	}
}

// 合并左式堆的实际例程
func Merge1(H1, H2 PriorityQueue) PriorityQueue {
	// 只有一个根节点而没有子节点的情况
	if H1.left == nil {
		H1.left = H2
	} else {
		// 合并总是向右子树上合并
		H1.right = Merge(H1.right, H2)
		if H1.left.npl < H1.right.npl {
			tmp := H1.left
			H1.left = H1.right
			H1.right = tmp
		}

		H1.npl = H1.right.npl + 1
	}

	return H1
}

// 左式堆插入例程
func Insert(X ElementType, H PriorityQueue) PriorityQueue {
	single := new(TreeNode)

	single.element = X
	single.npl = 0
	single.left = nil
	single.right = nil
	// 插入就是想原堆中合并新节点的过程
	H = Merge(single, H)

	return H
}

// 左式堆是满足堆序性质的
func DeleteMin(H PriorityQueue) (PriorityQueue, ElementType) {
	if H == nil {
		panic("PriorityQueue is empty")
	}

	LeftHeap := H.left
	RightHeap := H.right
	value := H.element
	H = nil

	return Merge(LeftHeap, RightHeap), value
}
