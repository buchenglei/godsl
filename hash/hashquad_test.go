package hash

import (
	//"fmt"
	"testing"
)

func Test_HashTable(t *testing.T) {
	table, err := NewHashTable(15, Hash)
	if err == nil && table.tableSize == 15 && table.theCells != nil {
		t.Log("  NewHashTable success!")
	} else {
		t.Error("  NewHashTable error!")
	}

	table.Insert("hello")
	old_pos := table.Find("hello")
	t.Log("  Please check position of the KEY! ", old_pos)

	table.Insert("hello")
	new_pos := table.Find("hello")
	if new_pos == old_pos {
		t.Log("  Insert success!")
	} else {
		t.Error("  Insert error!")
	}

	value := table.GetValue(new_pos)
	if table.GetValue(old_pos) == value && value == "hello" {
		t.Log("  GetValue success!")
	} else {
		t.Error("  GetValue error!")
	}

	// 清空hash表中的所有数据
	table.Empty()
	if table.GetValue(new_pos) == nil {
		t.Log("  Empty success!")
	} else {
		t.Error("  Empty error!")
	}


}

func Hash(str interface{}, tableSize int) int {
	var HashVal int = 0
	var c []byte

	if strings, ok := str.(string); ok {
		c = []byte(strings)
	}

	for _, v := range c {
		HashVal = (HashVal << 5) + int(v)
	}
	return HashVal % tableSize
}
