package datastructures

import (
	"fmt"
	"testing"
)

func Test_CQueueInit(t *testing.T) {
	cq := CQueueInit()
	fmt.Println(cq)
}

func Test_CQueueEnQueue_CQueueDeQueue(t *testing.T) {
	cq := CQueueInit()
	cq.CQueueEnQueue(1)
	fmt.Println(cq)

	cq.CQueueEnQueue(2)
	fmt.Println(cq)

	cq.CQueueEnQueue(3)
	fmt.Println(cq)

	cq.CQueueEnQueue(4)
	fmt.Println(cq)

	cq.CQueueEnQueue(5)
	fmt.Println(cq)

	fmt.Println(cq.CQueueDeQueue())
	fmt.Println(cq)

	fmt.Println(cq.CQueueDeQueue())
	fmt.Println(cq)

	fmt.Println(cq.CQueueDeQueue())
	fmt.Println(cq)

	fmt.Println(cq.CQueueDeQueue())
	fmt.Println(cq)

	fmt.Println(cq.CQueueDeQueue())
	fmt.Println(cq)

	fmt.Println(cq.CQueueDeQueue())
	fmt.Println(cq)

	//再入队
	cq.CQueueEnQueue(9)
	fmt.Println(cq)
	cq.CQueueEnQueue(8)
	fmt.Println(cq)
	cq.CQueueEnQueue(7)
	fmt.Println(cq)
	cq.CQueueEnQueue(6)
	fmt.Println(cq)

	cq.CQueueDeQueue()

	cq.CQueueEnQueue(5)
	fmt.Println(cq)
}
