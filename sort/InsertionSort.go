// 排序函数默认从小到大排列
package main

func InsertionSort(A []ElementType) {
	length := len(A)
	var j int

	for p := 1; p < length; p++ {
		Tmp := A[p]

		for j = p; j > 0 && A[j-1] > Tmp; j-- {
			// 将Tmp作为空穴，通过比较不断的向前移动
			// 避免了多次的元素交换的操作
			A[j] = A[j-1]
		}

		A[j] = Tmp
	}
}
