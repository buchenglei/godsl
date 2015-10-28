package main

// 归并排序的驱动程序
func msort(A []ElementType, TmpArray []ElementType,
	Left, Right int) {
	var Center int

	if Left < Right {
		Center = (Left + Right) / 2
		msort(A, TmpArray, Left, Center)
		msort(A, TmpArray, Center+1, Right)
		merge(A, TmpArray, Left, Center+1, Right)
	}
}

func merge(A []ElementType, TmpArray []ElementType,
	Lpos, Rpos, RightEnd int) {
	var LeftEnd, NumElements, TmpPos int

	LeftEnd = Rpos - 1
	TmpPos = Lpos
	NumElements = RightEnd - Lpos + 1

	// main loop
	// 比较数组的左半部分和右半部份
	// 将较小的值拷贝到TmpArray中去
	for Lpos <= LeftEnd && Rpos <= RightEnd {
		if A[Lpos] <= A[Rpos] {
			TmpArray[TmpPos] = A[Lpos]
			TmpPos++
			Lpos++
		} else {
			TmpArray[TmpPos] = A[Rpos]
			TmpPos++
			Rpos++
		}
	}

	// 将数组的左边剩余的部分拷到TmpArray中去
	for Lpos <= LeftEnd {
		TmpArray[TmpPos] = A[Lpos]
		TmpPos++
		Lpos++
	}
	// 将数组的右边剩余的部分拷到TmpArray中去
	for Rpos <= RightEnd {
		TmpArray[TmpPos] = A[Rpos]
		TmpPos++
		Rpos++
	}

	// Copy TmpArry back
	for i := 0; i < NumElements; i++ {
		A[RightEnd] = TmpArray[RightEnd]
		RightEnd--
	}
}

// 归并排序的主程序
func Mergesort(A []ElementType) {
	N := len(A)

	TmpArray := make([]ElementType, N)
	msort(A, TmpArray, 0, N-1)
}
