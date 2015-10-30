package main

import "fmt"

const Cutoff = 10

// 快速排序入口程序
func Quicksort(A []ElementType) {
	N := len(A)
	qsort(A, 0, N-1)
}

func swap(A []ElementType, Left, Right int) {
	tmp := A[Left]
	A[Left] = A[Right]
	A[Right] = tmp
}

// 实现三数中值方法的程序
func median3(A []ElementType, Left, Right int) ElementType {
	Center := (Left + Right) / 2

	// 将Left Center Right位置的值按照大小顺序排序
	// 使得A[Left] <= A[Center] <= A[Right]
	if A[Left] > A[Center] {
		swap(A, Left, Center)
	}
	if A[Left] > A[Right] {
		swap(A, Left, Right)
	}
	if A[Center] > A[Right] {
		swap(A, Center, Right)
	}

	// Right位置的值是肯定大于Center位置的值
	swap(A, Center, Right-1)

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
			for {
				i++
				if A[i] < Pivot {
				} else {
					break
				}
			}
			for {
				j--
				if A[j] > Pivot {
				} else {
					break
				}
			}
			if i < j {
				swap(A, i, j)
			} else {
				break
			}
		}
		swap(A, i, Right-1)

		qsort(A, Left, i-1)
		qsort(A, i+1, Right)
	} else {
		// 插入排序
		fmt.Println("InsertionSort")
		tmp := A[Left+1 : Right+1]
		InsertionSort(tmp)
	}

}

/*func insertionsort(A []ElementType, Left, Right int) {
	length := Right - Left + 1
	for p := left; p <= length; p++ {
		Tmp := A[p]

		for j = p; j > 0 && A[j-1] > Tmp; j-- {
			// 将Tmp作为空穴，通过比较不断的向前移动
			// 避免了多次的元素交换的操作
			A[j] = A[j-1]
		}

		A[j] = Tmp
	}
}
*/
