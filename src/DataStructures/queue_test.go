package datastructures

import (
	"fmt"
	"testing"
)

func Test_linkQueueInit(t *testing.T) {
	q := LinkQueueInit(5)
	fmt.Println(q)
}

func Test_EnQueue(t *testing.T) {
	q := LinkQueueInit(5)
	fmt.Println(q)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	fmt.Println(q)
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
}

func Test_GetHead(t *testing.T) {
	q := LinkQueueInit(5)
	fmt.Println(q)
	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	fmt.Println(q.GetHead())
	fmt.Println(q.DeQueue())
	fmt.Println(q.GetHead())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.GetHead())
	fmt.Println(q.DeQueue())
	fmt.Println(q.DeQueue())
	fmt.Println(q.GetHead())
}
