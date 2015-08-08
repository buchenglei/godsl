// 循环双向链表的实现 List
//
// Usage：
//	import "godsl"
//	func main() {
//		list := godsl.NewList()
//	}
package basic

import (
	"fmt"
)

// 声明链表中的节点结构
type Node struct {
	id   int
	data interface{}
	next *Node
	prev *Node
}

// 用于存放整个链表的结构
type List struct {
	header  *Node // 指向头节点的指针
	current *Node // 指向当前正在处理的节点的指针，处理完成后会将其设为nil
	count   int   //统计链表中节点的个数
	maxID   int   //用于设置每一个节点的唯ID
}

//返回当前节点保存的数据
func (n *Node) GetData() interface{} { return n.data }

//返回下一个节点
func (n *Node) GetNext() *Node { return n.next }

//返回前一个节点
func (n *Node) GetPrev() *Node { return n.prev }

// 创建一个新的链表，可以传递一个字符串作为链表的名称
// 返回一个指向初始化完成的链表的指针
// list := godsl.NewList()
func NewList() *List {
	list := new(List)
	list.header = new(Node)
	list.current = nil
	list.count = 0
	list.maxID = 0
	list.header.id = 0
	list.header.next = nil
	list.header.prev = nil

	return list
}

//创建一个新的节点
//该操作无需用户手动调用
func (l *List) createNewNode(data interface{}) {
	node := new(Node)

	//设置链表信息
	l.count += 1
	l.maxID += 1
	node.id = l.maxID

	//初始化新建节点
	node.data = data
	node.next = nil
	node.prev = nil

	//存储新建节点
	l.current = node
}

//检查节点是否为空
func (l *List) IsEmpty() bool { return l.header.next == nil }

//返回链表的长度
func (l *List) Len() int { return l.count }

func (l *List) setPrevAndNext(prev *Node, next *Node) {
	//当链表为空的时候
	//处理插入到头节点之前或之后的情况
	if prev == nil {
		l.current.next = next
		l.current.prev = next
		next.prev = l.current
		next.next = l.current
	}
	if next == nil {
		l.current.next = prev
		l.current.prev = prev
		prev.next = l.current
		prev.prev = l.current
	} else {
		l.current.next = next
		l.current.prev = prev
		next.prev = l.current
		prev.next = l.current
	}
}

//将节点追加到链表的末尾
func (l *List) Append(data interface{}) int {
	//创建一个新的节点，存储到list.current中
	l.createNewNode(data)

	//将获得的节点追加到链表的尾部
	var tail *Node
	if l.header.prev == nil {
		tail = l.header
	} else {
		tail = l.header.prev
	}
	//当前的p指向最后一个节点
	l.setPrevAndNext(tail, tail.next)
	//将新创建节点的id返回
	return l.current.id

}

//将节点插入到指定ID的节点之后
//插入成功，则返回刚刚插入节点的id
func (l *List) InsertAfterID(id int, data interface{}) (int, error) {
	// 用于存储待插入的节点
	finded, err := l.FindByID(id)
	if err != nil {
		return -1, err
	}

	//创建一个新的节点
	l.createNewNode(data)
	l.setPrevAndNext(finded, finded.next)

	return l.current.id, nil

}

//插入到指定ID的节点之前
//插入成功，则返回刚刚插入节点的id
func (l *List) InsertBeforeID(id int, data interface{}) (int, error) {
	finded, err := l.FindByID(id)
	if err != nil {
		return -1, err
	}
	l.createNewNode(data)
	l.setPrevAndNext(finded.prev, finded)

	return l.current.id, nil

}

//按照ID查找相应的节点并返回
func (l *List) FindByID(id int) (*Node, error) {
	if id == 0 {
		return l.header, nil
	}
	p := l.header.next
	for p.id != 0 {
		if id == p.id {
			return p, nil
		} else {
			p = p.next
		}
	}
	return nil, fmt.Errorf("can't find this id %d", id)
}

//按节点内容查找
func (l *List) FindByData(data interface{}, count int) ([]*Node, error) {
	tmp := l.header.next
	nodes := make([]*Node, count)
	i := 0
	//根据count的值来设置slice的大小
	for tmp.id != 0 && count != 0 {
		if tmp.data == data {
			count -= 1
			nodes[i] = tmp
			i++
		}
		tmp = tmp.next
	}

	return nodes, nil
}

//删除指定id的节点
func (l *List) DeleteByID(id int) error {
	deleted, err := l.FindByID(id)
	if err != nil {
		return err
	}
	deleted.prev = deleted.next
	l.count -= 1
	return nil
}

//删除指定内容的节点
func (l *List) DeleteByData(data interface{}, count int) error {
	deleted, err := l.FindByData(data, count)
	if err != nil {
		return err
	}
	for _, n := range deleted {
		if n != nil {
			n.prev = n.next
		}
	}

	l.count -= 1
	return nil
}

// 将链表自身打印出来
func (l *List) PrintMe() {
	tmp := l.header.next
	fmt.Printf("List Count:%d\n----------\n", l.count)

	for tmp.id != 0 {
		fmt.Println("id: ", tmp.id, "\tdata:\n\t ", tmp.data)
		tmp = tmp.next
	}
}
