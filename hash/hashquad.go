// 开放定址法实现的hash表
package hash

import "fmt"

// 定义hash表中每一项的类型
const (
	Legitimate   = iota // 已存在和合法数据
	Empty        = iota // 当前位置为空
	Deleted      = iota // 当前数据已被删除
	MinTableSize = 10   // hash表大小的最小值
)

// 自定义类型，接受一个hash函数
type HashFunc func(data interface{}, tableSize int) int

// 每一个hash表中的项目
type HashEntry struct {
	data interface{}
	kind int
}

// hash表结构
type HashTable struct {
	tableSize int
	theCells  []*HashEntry
	hash      HashFunc // 注册的hash函数
}

// 判断当前的树是否为素数，不是则返回下一个素数
func NextPrime(num int) int {
	// TODO
	return num
}

// 创建一个新的hash表
func NewHashTable(size int, hash HashFunc) (*HashTable, error) {
	if size < MinTableSize {
		return nil, fmt.Errorf("Table size too small!")
	}
	if hash == nil {
		return nil, fmt.Errorf("hashTable need a hash function!")
	}

	// 初始化一个新的hash表
	hashTable := new(HashTable)
	// 获取一个素数
	hashTable.tableSize = NextPrime(size)
	// 创建一个用于存储HashEntry的slice
	hashTable.theCells = make([]*HashEntry, size)
	// 注册hash函数
	hashTable.hash = hash

	// 初始化theCells
	for i := 0; i < hashTable.tableSize; i++ {
		hashTable.theCells[i] = new(HashEntry)
		hashTable.theCells[i].data = nil
		hashTable.theCells[i].kind = Empty
	}

	return hashTable, nil
}

func (t *HashTable) Find(data interface{}) int {
	var collisionNum int = 0
	// 获取data应该存放的位置
	currentPos := t.hash(data, t.tableSize)
	// 从当前位置寻找一个空的位置
	for t.theCells[currentPos].kind != Empty &&
		t.theCells[currentPos].data != data {

		// ***平方探测法寻找下一个存储单元***
		collisionNum += 1
		currentPos += 2*collisionNum - 1
		// 如果当前位置超过了hash表的大小
		// 则将currentPos放在hash开头的位置寻找
		if currentPos >= t.tableSize {
			currentPos -= t.tableSize
		}
	}

	// 返回查找到的位置
	// 可能是一个已存在HashEntry的位置
	// 也可能是一个空位置或是被删除的位置
	return currentPos
}

func (t *HashTable) Insert(data interface{}) {
	// finded is a pointer to HashEntry
	position := t.Find(data)
	entry := t.theCells[position]
	if entry.kind != Legitimate {
		entry.kind = Legitimate
		entry.data = data
	}

}

// 这里的清空并不会删除原有的数据
// 而是将每一个HashEntry的kind置为Delete
func (t *HashTable) Empty() {
	for i := 0; i < t.tableSize; i++ {
		if t.theCells[i] == nil {
			continue
		}
		t.theCells[i].kind = Deleted
	}
}

// 根据Find到的index，判断相应的值是否存在
func (t *HashTable) GetValue(index int) interface{} {
	if index > t.tableSize {
		return nil
	}

	entry := t.theCells[index]

	if entry.kind == Legitimate {
		return entry.data
	} else {
		return nil
	}
}
