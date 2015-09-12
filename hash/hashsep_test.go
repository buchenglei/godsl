package hash

import (
	"testing"
)

func Test_Hashsep(t *testing.T) {
	H, err := NewHash_1(10, Hash)
	if err != nil {
		t.Error(err)
	}

	H.Insert("hello world")
	node := H.Find("hello world")
	if node.GetData() == "hello world" {
		t.Log("  Insert and Find success!")
	} else {
		t.Error("  Insertand Find error!")
	}

	H.Delete("hello world")
	node = H.Find("hello world")
	if node == nil {
		t.Log("  Delete success!")
	} else {
		t.Error("  Delete error!")
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
