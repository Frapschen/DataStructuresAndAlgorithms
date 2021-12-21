package datastructures

import (
	"fmt"
	"testing"
)

func Test_dLinkListInit(t *testing.T) {
	dlink := DLinkListInit()
	fmt.Println(dlink)
}

func Test_adddNode(t *testing.T) {
	dlink := DLinkListInit()
	dlink.AddDNode(1)
	dlink.AddDNode(2)
	dlink.AddDNode(3)
	dlink.AddDNode(4)
	dlink.AddDNode(5)
	a, b := dlink.DLinkListToSlice()
	fmt.Println(a, b)
}

func Test_DelDLinkList(t *testing.T) {
	dlink := DLinkListInit()
	dlink.AddDNode(1)
	dlink.AddDNode(2)
	dlink.AddDNode(3)
	dlink.AddDNode(4)
	dlink.AddDNode(5)
	a, b := dlink.DLinkListToSlice()
	fmt.Println(a, b)

	dlink.DelDLinkList(2)
	dlink.DelDLinkList(1)
	fmt.Println(dlink.DLinkListToSlice())
}

func Test_FindDLinkList(t *testing.T) {
	dlink := DLinkListInit()
	dlink.AddDNode(1)
	dlink.AddDNode(2)
	dlink.AddDNode(3)
	dlink.AddDNode(4)
	dlink.AddDNode(5)
	fmt.Println(dlink.FindDLinkList(9))
}
