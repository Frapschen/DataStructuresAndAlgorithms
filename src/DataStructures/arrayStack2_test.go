package datastructures

import (
	"fmt"
	"testing"
)

func TestInitStack2(t *testing.T) {
	s := InitStack2()
	fmt.Println(s)
}

func TestPush2(t *testing.T) {
	s := InitStack2()
	s.Push2(1)
	s.Push2(2)
	s.Push2(3)
	s.Push2(4)
	s.Push2(5)
	s.Push2(6)
	fmt.Println(s)
}

func TestPop2(t *testing.T) {
	s := InitStack2()
	s.Push2(1)
	s.Push2(2)
	s.Push2(3)
	s.Push2(4)
	s.Push2(5)
	s.Push2(6)
	fmt.Println(s)
	s.Pop2()
	fmt.Println(s.GetTop2())
	s.Pop2()
	s.Pop2()
	fmt.Println(s.GetTop2())
}
