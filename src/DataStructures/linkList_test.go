package datastructures

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	ll := Init()
	fmt.Println(ll)
}
func TestAdd(t *testing.T) {
	ll := Init()
	fmt.Println(ll.ToSlice())
	ll.AddNode(1)
	ll.AddNode(2)
	fmt.Println(ll.ToSlice())
	ll.AddNode(3)
	ll.AddNode(4)
	fmt.Println(ll.ToSlice())
}

func TestDel(t *testing.T) {
	ll := Init()
	ll.AddNode(1)
	ll.AddNode(2)
	ll.AddNode(3)
	ll.AddNode(4)
	fmt.Println("删除前：", ll.ToSlice())
	ll.DelNode(3)
	fmt.Println("删除后：", ll.ToSlice())
}

func TestFind(t *testing.T) {
	ll := Init()
	ll.AddNode(1)
	ll.AddNode(2)
	ll.AddNode(3)
	ll.AddNode(4)
	fmt.Println("data：", ll.ToSlice())
	fmt.Println("查询结果：", ll.Find(4))
}
