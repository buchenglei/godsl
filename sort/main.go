// 用于测试各个排序函数的正确性和效率
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ElementType int

// 当前正在调用的排序方法
var currentOrderName string = ""

// 待测试数据
var TestData []ElementType

// 记录算法开始的时间
var startTime time.Time

func main() {
	TestData = CreateTestData(100)

	// 插入排序
	currentOrderName = "插入排序"
	start()
	InsertionSort(TestData)
	CheckOrder()
	end()

}

func CreateTestData(length int) []ElementType {
	if length <= 0 {
		panic("No test data!")
	}

	data := make([]ElementType, length)
	for i := 0; i < length; i++ {
		data[i] = ElementType(rand.Intn(200))
	}

	return data
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

// 计算排序算法的运行时间
func start() {
	startTime = time.Now()
}

func end() {
	endTime := time.Now()
	nanosecond := float32(endTime.Nanosecond() - startTime.Nanosecond())
	millisecond := nanosecond / float32(time.Millisecond)
	microsecond := nanosecond / float32(time.Microsecond)
	second := nanosecond / float32(time.Second)

	fmt.Printf("共耗时 %.2f 秒，%.2f 毫秒，%.2f 微秒, %.2f 纳秒\n",
		second, millisecond, microsecond, nanosecond)

	fmt.Println()
}
