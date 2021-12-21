package datastructures

import (
	"fmt"
	"testing"
)

func Test_Order(t *testing.T) {
	b := RondomBiTree()
	PreOrder(&b)
	fmt.Println()
	InOrder(&b)
	fmt.Println()
	PostOrder(&b)
	fmt.Println()
}

func Test_Order2(t *testing.T) {
	b := RondomBiTree()
	PreOrder2(&b)
	fmt.Println()
	InOrder2(&b)
	fmt.Println()
	PostOrder2(&b)
}
