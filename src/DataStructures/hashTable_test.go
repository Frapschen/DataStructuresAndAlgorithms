package datastructures

import (
	"fmt"
	"testing"
)

func Test_hashTable(t *testing.T) {
	table := &hashTbale{}
	table.HashTableAdd("AA", 1)
	table.HashTableAdd("BB", 2)
	table.HashTableAdd("CC", 3)
	table.HashTableAdd("DD", 4)
	fmt.Println(table.data)
	fmt.Println(table.HashTableGet("AA"))
	fmt.Println(table.HashTableGet("BB"))
	fmt.Println(table.HashTableGet("CC"))
	fmt.Println(table.HashTableGet("DD"))

	fmt.Println(table.HashTableRemove("BB"))
	fmt.Println(table.HashTableRemove("CC"))
	fmt.Println(table.data)
}
