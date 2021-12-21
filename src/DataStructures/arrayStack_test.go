package datastructures

import (
	"fmt"
	"testing"
)

func TestInitStack(t *testing.T) {
	s := InitStack()
	fmt.Println(s)
}

func TestPush(t *testing.T) {
	s := InitStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	fmt.Println(s)
}

func TestPop(t *testing.T) {
	s := InitStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s.GetTop())
	s.Pop()
	s.Pop()
	fmt.Println(s.GetTop())
}
