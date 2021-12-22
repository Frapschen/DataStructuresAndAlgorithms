package search

import "fmt"

type avlNode struct {
	data, high     int
	lchild, rchild *avlNode
}

func BuildAVL(num []int) *avlNode {
	if len(num) == 0 {
		return nil
	}
	root := &avlNode{data: num[0]}
	for i := 1; i < len(num); i++ {
		InsertAVL(root, num[i])
	}
	return root
}

func InsertAVL(avl *avlNode, num int) int {

	return 1
}

//中序
func InOrderAVL(b *avlNode) {
	if b != nil {
		InOrderAVL(b.lchild)
		fmt.Print(b.data, " ")
		InOrderAVL(b.rchild)
	}
}
