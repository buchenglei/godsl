package hash

import "fmt"

type listNode struct {
	data interface{}
	next *listNode
}

type hashFunc func(data interface{}) int

// 分离链接散列表的类型声明
type HashTbl_1 struct {
	hash      hashFunc
	tableSize int
	lists     []*listNode
}

// 定义hash表大小的最小值
const minTableSize = 5

func newNode(data interface{}) *listNode {
	n := new(listNode)
	n.data = data
	n.next = nil
	return n
}

// 创建一个新的分离链接的散列表
func NewHash_1(tableSize int, hash hashFunc) (*HashTbl_1, error) {
	if tableSize < minTableSize {
		return nil, fmt.Errorf("Table size too small!")
	}

	if hash == nil {
		return nil, fmt.Errorf("Need a Hash function!")
	}

	// 创建一个新的hash表
	H := new(HashTbl_1)
	H.tableSize = tableSize
	// 注册hash函数
	H.hash = hash
	// 用于存储链表的slice
	H.lists = make([]*listNode, tableSize)
	// 初始化slice中每个元素链表的头结点
	for i := 0; i < tableSize; i++ {
		H.lists[i] = newNode(nil)
	}

	return H, nil
}

// 查找指定的关键字是否在hash表中
func (h *HashTbl_1) Find(key interface{}) *listNode {
	L := h.lists[h.hash(key)]
	P := L.next // 头结点不保存数据
	for P != nil && P.data != key {
		P = P.next
	}

	return P
}

// 向hash表中插入新的数据
func (h *HashTbl_1) Insert(key interface{}) {
	finded := h.Find(key)

	if finded == nil {
		// 计算待存储的位置，并返回该位置上的链表的头结点
		L := h.lists[h.hash(key)]
		// 将待插入的关键字插入链表的头结点
		newnode := newNode(key)
		newnode.next = L.next
		L.next = newnode
	}
}

// 删除hash表中指定的关键字
func (h *HashTbl_1) Delete(key interface{}) {
	L := h.lists[h.hash(key)]

	for L.next != nil && L.next.data != key {
		L = L.next
	}

	L.next = L.next.next

}
