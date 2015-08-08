package basic

import (
	"fmt"
	"testing"
)

func Test_List(t *testing.T) {
	//定义待测试的数据
	var testData [10]string
	//测试使用的变量
	var (
		id      int
		err     error
		isempty bool
		node    *Node
		nodes   []*Node
	)

	for i := 0; i < 10; i++ {
		testData[i] = fmt.Sprintf("testing......%d", i+1)
	}

	list := NewList()

	//判断新建的链表是否为空
	isempty = list.IsEmpty()
	if isempty {
		t.Log(" When list is new, IsEmpty Success!")
	} else {
		t.Error(" IsEmpty should return true, but not, please check")
	}

	//追加节点
	for i := 0; i < 10; i++ {
		id = list.Append(testData[i])
		if id == i+1 {
			t.Log(" Append Success!")
		}
	}

	//插入到头节点之后
	id, err = list.InsertAfterID(0, "testing......11")
	if id == 11 && err == nil {
		t.Log(" Id = 11, InsertAfterID Success!")
	} else {
		t.Error(" Except id = 11 and err = nil, but not")
	}

	//插入到尾节点之后
	id, err = list.InsertAfterID(10, "testing......12")
	if id == 12 && err == nil {
		t.Log(" Id = 12, InsertAfterID Success!")
	} else {
		t.Error(" Except id = 12 and err = nil, but not")
	}

	//插入到头节点之前
	id, err = list.InsertBeforeID(0, "testing......13")
	if id == 13 && err == nil {
		t.Log(" Id = 13, InsertBeforeID Success!")
	} else {
		t.Error(" Except id = 13 and err = nil, but not")
	}
	//插入到头节点之前
	id, err = list.InsertBeforeID(12, "testing......14")
	if id == 14 && err == nil {
		t.Log(" Id = 14, InsertBeforeID Success!")
	} else {
		t.Error(" Except id = 14 and err = nil, but not")
	}

	//按id查找
	node, err = list.FindByID(1)
	if node.id == 1 && err == nil {
		t.Log(" Id = 1, FindByID Success!")
	} else {
		t.Error(" Except id = 1 and err = nil, but not")
	}

	fmt.Println("list.Len() is ", list.Len())
	fmt.Println("node.id=1 is ", node.GetData())

	//测试查找不存在的节点
	node, err = list.FindByID(100)
	fmt.Println(err)
	if node == nil && err != nil {
		t.Log(" FindByID Success!")
	} else {
		t.Error(" Find a unexsits id")
	}

	//按数据查找
	nodes, err = list.FindByData("testing......14", 1)
	if nodes[0].id == 14 && err == nil {
		t.Log(" FindByData Success")
	} else {
		t.Error(" nodes[0].id should be 14, but not")
	}

	list.PrintMe()
}
