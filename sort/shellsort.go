// 希尔排序
package main

func Shellsort(A []ElementType) {
	N := len(A)
	var j int

	for Increment := N / 2; Increment > 0; Increment /= 2 {
		for i := Increment; i < N; i++ {
			Tmp := A[i]
			for j = i; j >= Increment; j -= Increment {
				if Tmp < A[j-Increment] {
					A[j] = A[j-Increment]
				} else {
					break
				}
			}
			A[j] = Tmp
		}
	}
}
