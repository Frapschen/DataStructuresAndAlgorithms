package search

import (
	"fmt"
)

type bstTNode struct {
	data           int
	lchild, rchild *bstTNode
}

func BuildBST(num []int) *bstTNode {
	if len(num) == 0 {
		return nil
	}
	root := &bstTNode{data: num[0]}
	for i := 1; i < len(num); i++ {
		InsertBST(root, num[i])
	}
	return root
}

//插入
func InsertBST(b *bstTNode, num int) int {
	if b.data != num {
		if b.data > num {
			if b.lchild != nil {
				InsertBST(b.lchild, num)
			} else {
				b.lchild = &bstTNode{data: num}
				return 1
			}
		} else {
			if b.rchild != nil {
				InsertBST(b.rchild, num)
			} else {
				b.rchild = &bstTNode{data: num}
				return 1
			}
		}
	}
	return -1
}

//删除
func DelBST(b *bstTNode, n int) bool {
	parent := b
	flag := 0 //0代表左孩子，1代表有孩子
	for b != nil && b.data != n {
		if b.data > n {
			parent = b
			b = b.lchild
			flag = 0
		} else {
			parent = b
			b = b.rchild
			flag = 1
		}
	}
	if b == nil {
		return false
	}
	//是叶子节点
	if b.rchild == nil && b.lchild == nil {
		if flag == 0 {
			parent.lchild = nil
		} else {
			parent.rchild = nil
		}
	} else if b.lchild != nil && b.rchild == nil { //只有左孩子
		b.data = b.lchild.data
		b.lchild = nil
	} else if b.rchild != nil && b.lchild == nil { //只有右孩子
		b.data = b.rchild.data
		b.rchild = nil
	} else { //左右孩子有都
		if b.rchild.rchild == nil { //右孩子为叶子节点的可以简单处理
			b.data = b.rchild.data
			b.rchild = nil
			return true
		}
		inOrderNode := rightTreeInOrder(b.rchild) //右子树中序遍历的第一个节点
		b.data = inOrderNode.data
		DelBST(b.rchild, inOrderNode.data) //在这个右子树中删除找到的中序节点
	}
	return true
}

//中序遍历的第一个节点
func rightTreeInOrder(b *bstTNode) *bstTNode {
	var target *bstTNode
	for b.lchild != nil {
		target = b.lchild
		b = b.lchild
	}
	return target
}

//查询
func FindBST(b *bstTNode, n int) bool {
	for b != nil && b.data != n {
		if b.data > n {
			b = b.lchild
		} else {
			b = b.rchild
		}
	}
	if b == nil {
		return false
	}
	return true
}

//中序
func InOrder(b *bstTNode) {
	if b != nil {
		InOrder(b.lchild)
		fmt.Print(b.data, " ")
		InOrder(b.rchild)
	}
}
