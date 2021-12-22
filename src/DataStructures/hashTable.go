package datastructures

import (
	"hash/fnv"
)

//哈希表原理很简单，用哈希算法算出位置，用拉链发解决冲突(如果有)
/*
https://github.com/Frapschen/Data-Structures-and-Algorithms/blob/master/HashTable/HashTable.go
*/

const MAXLENGTH uint = 5

type tableItem struct {
	key   string
	value int
	next  *tableItem
}

type hashTbale struct {
	data [MAXLENGTH]*tableItem
}

func (t *hashTbale) HashTableAdd(key string, value int) {
	position := generateHash(key)
	if t.data[position] == nil {
		t.data[position] = &tableItem{key: key, value: value}
		return
	}
	temp := t.data[position]
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = &tableItem{key: key, value: value}
	return
}

func (t *hashTbale) HashTableGet(key string) (int, bool) {
	position := generateHash(key)
	temp := t.data[position]
	for temp != nil {
		if temp.key == key {
			return temp.value, true
		}
		temp = temp.next
	}
	return 0, false
}

func (t *hashTbale) HashTableSet(key string, value int) bool {
	position := generateHash(key)
	temp := t.data[position]
	for temp != nil {
		if temp.key == key {
			temp.value = value
			return true
		}
		temp = temp.next
	}
	return false
}

func (t *hashTbale) HashTableRemove(key string) bool {
	position := generateHash(key)
	if t.data[position] == nil {
		return false
	}
	if t.data[position].key == key {
		t.data[position] = t.data[position].next
		return true
	}
	temp := t.data[position]
	for temp.next != nil {
		if temp.next.key == key {
			temp.next = temp.next.next
			return true
		}
		temp = temp.next
	}
	return false
}

//产生hash值
func generateHash(key string) uint {
	hash := fnv.New32a()
	hash.Write([]byte(key))
	return (uint(hash.Sum32()) % MAXLENGTH)
}
