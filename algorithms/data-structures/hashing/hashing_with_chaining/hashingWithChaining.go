package hashing_with_chaining

import (
	"fmt"
	"math"
	"ultimate-go/algorithms/data-structures/lists/singlylinkedlist"
)

/*
https://en.wikipedia.org/wiki/Hash_table
https://en.wikipedia.org/wiki/Associative_array#Hash_table_implementations


ciphers: 加密
hash: 本质是digest即摘要
hmac: 本质是digest即摘要
sign： 签名


hashtable， 就是一个存key, value的Associative array
hashtable, hashFunction, 通过hashFunction算法计算出一个index
hashtable, key ， 即index， 之后发value存入index对应的linked list
hashtable, collision resolution, 为什么存linked list, 因为计算出的index可以相同， 即

add, insert(key, value interface{})
remove, delete(key, value interface{})
lookup/find/search(key)
change

all entries, O(n)
all keys, O(n)
all values, O(n)

size, O(1)
isEmpty, O(1)
isFull, O(1)
loadFactor, O(1)
*/
const arrayLength = 10

type HashTable struct {
	data [arrayLength]*singlylinkedlist.LinkedList
}
type listData struct {
	key   string
	value interface{}
}

// 计算hash值
func Hash(s string) int {
	h := 0
	for pos, char := range s {
		h += int(char) * int(math.Pow(31, float64(len(s)-pos+1)))
	}
	return h
}

// 计算table中的index
func Index(hash int) int {
	return hash % arrayLength
}

func (h *HashTable) Add(key string, value interface{}) *HashTable {
	index := Index(Hash(key))
	// index不存在的情况下， 新建一个linked list
	if h.data[index] == nil {
		h.data[index] = new(singlylinkedlist.LinkedList)
		h.data[index].Append(listData{
			key:   key,
			value: value,
		})
	} else { // index存在的情况下, 不需要新建linked list
		node := h.data[index].Head
		for {
			// key 存在的情况下， 覆盖
			if node != nil {
				d := node.Value.(listData)
				if d.key == key {
					d.value = value
					break
				}
			} else { // key不存在的情况下， 添加一个node
				h.data[index].Append(listData{
					key:   key,
					value: value,
				})
			}
		}

	}
	return h
}
func (h *HashTable) Get(k string) (value interface{}, ok bool) {
	index := Index(Hash(k))
	if h.data[index] == nil {
		return nil, false
	}
	node := h.data[index].Head
	if node != nil {
		d := node.Value.(listData)
		if d.key == k {
			fmt.Print("index: ", index, ", ")
			h.data[index].Print()
			return d.value, true
		}
	} else {
		return nil, false
	}
	return nil, false
}
