// 用于测试各个排序函数的正确性和效率
package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type ElementType int

// 当前正在调用的排序方法
var currentOrderName string = ""

// 待测试数据
var TestData []ElementType

// 测试数据的个数， 默认 1000 个
var Numbers int = 1000

func main() {

	if len(os.Args) == 2 {
		t1, err := strconv.ParseInt(os.Args[1], 10, 32)
		if t1 > 0 && err == nil {
			Numbers = int(t1)
		}
	}
	// 初始化测试slice
	TestData = make([]ElementType, Numbers)

	// 插入排序
	Test(InsertionSort, "插入排序")
	// 希尔排序
	Test(Shellsort, "希尔排序")
	// 堆排序
	Test(Heapsort, "堆排序  ")
	// 归并排序
	Test(Mergesort, "归并排序")
	// 快速排序
	Test(Quicksort, "快速排序")

}

// 统一的测试函数
func Test(f func(A []ElementType), name string) {
	InitTestData()
	currentOrderName = name
	start := time.Now()
	f(TestData)
	CheckOrder()
	end(start)

}

func InitTestData() {
	if Numbers <= 0 {
		panic("No test data!")
	}

	for i := 0; i < Numbers; i++ {
		TestData[i] = ElementType(rand.Intn(Numbers))
	}

}

func CheckOrder() {
	// 检查数组中的元素是否是从大到小排列的
	n := len(TestData)
	var s bool = true
	for i := 0; i < n-1; i++ {
		// 前一个元素比后一个元素大，出错
		if TestData[i] > TestData[i+1] {
			fmt.Printf("在 %s 排序中，第 %d 个元素和第 %d 个元素关系错误\n",
				currentOrderName, i, i+1)
			s = false
		}
	}

	if s {
		fmt.Print(currentOrderName, "\t排序正确\t")
	} else {
		fmt.Print("\nError: ", currentOrderName, "\t排序错误\t")
	}

}

func end(startTime time.Time) {
	endTime := time.Now()
	nanosecond := float64(endTime.Nanosecond() - startTime.Nanosecond())
	millisecond := nanosecond / float64(time.Millisecond)
	microsecond := nanosecond / float64(time.Microsecond)
	second := nanosecond / float64(time.Second)

	fmt.Printf("共耗时 %f 秒，%f 毫秒，%f 微秒, %f 纳秒\n",
		second, millisecond, microsecond, nanosecond)
}
