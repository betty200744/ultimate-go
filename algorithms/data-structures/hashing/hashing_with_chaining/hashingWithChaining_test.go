package hashing_with_chaining

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {

	m.Run()
}
func Test_Hash(t *testing.T) {
	t.Run("hash", func(t *testing.T) {
		fmt.Println(Hash("hello"))
	})
	t.Run("index", func(t *testing.T) {
		fmt.Println(Index(Hash("hello")))
	})
}
func TestHashTable(t *testing.T) {
	hashTable := new(HashTable)
	hashTable.Add("a", "a")
	hashTable.Add("b", "b")
	hashTable.Get("a")
	hashTable.Get("b")
}
