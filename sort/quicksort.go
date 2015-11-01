package main

// 设置一个临界点
// 当数组长度过小的时候，使用插入排序
const Cutoff = 10

// 快速排序入口程序
func Quicksort(A []ElementType) {
	N := len(A)
	qsort(A, 0, N-1)
}

// 实现三数中值方法的程序
func median3(A []ElementType, Left, Right int) ElementType {
	Center := (Left + Right) / 2

	// 将Left Center Right位置的值按照大小顺序排序
	// 使得A[Left] <= A[Center] <= A[Right]
	if A[Left] > A[Center] {
		A[Left], A[Center] = A[Center], A[Left]
	}
	if A[Left] > A[Right] {
		A[Left], A[Right] = A[Right], A[Left]
	}
	if A[Center] > A[Right] {
		A[Center], A[Right] = A[Right], A[Center]
	}

	// Right位置的值是肯定大于Center位置的值
	A[Center], A[Right-1] = A[Right-1], A[Center]

	return A[Right-1]
}

// 快速排序主程序
func qsort(A []ElementType, Left, Right int) {
	// 如果元素个数过少，则使用插入排序提高效率
	if Left+Cutoff <= Right {
		Pivot := median3(A, Left, Right)
		i := Left
		j := Right - 1
		for {
			// 这里的写法比较恶心，翻译成C语言就是：
			// while(A[++i] < Pivot) {}
			for {
				if i += 1; A[i] < Pivot {
				} else {
					break
				}
			}
			/// while(A[--j] > Pivot) {}
			for {
				if j -= 1; A[j] > Pivot {
				} else {
					break
				}
			}
			if i < j {
				A[i], A[j] = A[j], A[i]
			} else {
				break
			}
		}
		A[i], A[Right-1] = A[Right-1], A[i]

		qsort(A, Left, i-1)
		qsort(A, i+1, Right)
	} else {
		// 插入排序
		// 这是一个容易出错的地方
		// 指明范围的时候a:b，包含a但不包含b
		tmp := A[Left : Right+1]
		Insertionsort(tmp)
	}

}
