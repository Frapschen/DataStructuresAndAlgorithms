package datastructures

import "fmt"

type biTNode struct {
	data           int
	lchild, rchild *biTNode
}

func RondomBiTree() biTNode {
	root := biTNode{data: 1}
	temp1 := &biTNode{data: 2}
	temp2 := &biTNode{data: 3}
	temp3 := &biTNode{data: 4}
	temp4 := &biTNode{data: 5}
	temp5 := &biTNode{data: 6}
	temp6 := &biTNode{data: 7}
	temp7 := &biTNode{data: 8}
	temp8 := &biTNode{data: 9}

	root.lchild = temp1
	temp1.rchild = temp2
	temp1.lchild = temp3
	temp3.rchild = temp4

	root.rchild = temp5
	temp5.rchild = temp6
	temp5.lchild = temp7
	temp7.lchild = temp8

	return root
}

//先序
func PreOrder(b *biTNode) {
	if b != nil {
		fmt.Print(b.data, " ")
		PreOrder(b.lchild)
		PreOrder(b.rchild)
	}
}

//中序
func InOrder(b *biTNode) {
	if b != nil {
		InOrder(b.lchild)
		fmt.Print(b.data, " ")
		InOrder(b.rchild)
	}
}

//后续
func PostOrder(b *biTNode) {
	if b != nil {
		PostOrder(b.lchild)
		PostOrder(b.rchild)
		fmt.Print(b.data, " ")
	}
}

//先序2
func PreOrder2(b *biTNode) {
	s := InitStack2()
	for b != nil || !s.IsEmpty() {
		if b != nil {
			fmt.Print(b.data, " ")
			s.Push2(b)
			b = b.lchild
		} else {
			b = s.Pop2().(*biTNode).rchild
		}
	}
}

//中序2
func InOrder2(b *biTNode) {
	s := InitStack2()
	for b != nil || !s.IsEmpty() {
		if b != nil {
			s.Push2(b)
			b = b.lchild
		} else {
			b = s.Pop2().(*biTNode)
			fmt.Print(b.data, " ")
			b = b.rchild
		}

	}
}

//后续2
func PostOrder2(b *biTNode) {
	s := InitStack2()
	var flag *biTNode
	for b != nil || !s.IsEmpty() {
		if b != nil {
			s.Push2(b)
			b = b.lchild
		} else if s.GetTop2().(*biTNode).rchild != nil &&
			s.GetTop2().(*biTNode).rchild != flag {
			b = s.GetTop2().(*biTNode).rchild
		} else {
			b = s.Pop2().(*biTNode)
			flag = b
			fmt.Print(b.data, " ")
			b = nil
		}
	}
}
