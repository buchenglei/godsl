package main

func leftchild(i int) int {
	return (2*i + 1)
}

// 构造二叉堆时，数组中Index=0并不特殊处理
func percdown(A []ElementType, i, N int) {
	var Child int
	var Tmp ElementType

	for Tmp = A[i]; leftchild(i) < N; i = Child {
		Child := leftchild(i)
		// 找出左右子节点中较小的那一个
		// Child != N - 1目的是防止Child指向最后一个元素
		// 从而防止数组越界的问题
		if Child != N-1 && A[Child+1] < A[Child] {
			Child++
		}
		// 从小到大排列
		if Tmp > A[Child] {
			A[i] = A[Child]
		} else {
			break
		}
		// 把Tmp放到它应有的位置
	}
	A[i] = Tmp

}

func Heapsort(A []ElementType) {
	N := len(A)

	// Build Heap
	for i := N / 2; i >= 0; i-- {
		percdown(A, i, N)
	}
	// Delete Min
	for i := N - 1; i > 0; i-- {
		tmp := A[0]
		A[0] = A[i]
		A[i] = tmp
		percdown(A, 0, i)
	}

}
