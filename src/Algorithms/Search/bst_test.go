package search

import (
	"testing"
)

func Test_bst(t *testing.T) {
	// data := []int{11, 2, 3, 5, 6, 7, 8, 9, 10, 13, 15}
	// bst := BuildBST(data)
	// InOrder(bst)
	// fmt.Println("查找:", FindBST(bst, 12))

	test_data := []int{53, 17, 78, 9, 45, 65, 87, 23}
	bst := BuildBST(test_data)
	InOrder(bst)
	println()

	//删除叶子
	// DelBST(bst, 23)
	// DelBST(bst, 87)
	// InOrder(bst)
	// println()

	//删除只有一颗左/右子树的节点
	// DelBST(bst, 45)
	// DelBST(bst, 23)
	// InOrder(bst)
	// println()

	//删除有两个孩子的节点
	//1.孩子是叶子节点的
	DelBST(bst, 78)
	InOrder(bst)
	println()
}
