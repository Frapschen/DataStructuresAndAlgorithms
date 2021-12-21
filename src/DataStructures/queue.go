package datastructures

type queueNode struct {
	data int
	next *queueNode
}

type linkQueue struct {
	head, rear *queueNode
	size       int
	maxSize    int
}

//初始化
func LinkQueueInit(max int) linkQueue {
	head := &queueNode{}
	return linkQueue{size: 0,
		maxSize: max,
		head:    head,
		rear:    head}
}

//入队
func (q *linkQueue) EnQueue(data int) bool {
	if q.size == q.maxSize {
		return false
	}
	temp := &queueNode{data: data,
		next: nil}
	if q.rear == q.head {
		q.head.next = temp
		q.rear = temp
		q.size++
	} else {
		q.rear.next = temp
		q.rear = temp
		q.size++
	}
	return true
}

//出队
func (q *linkQueue) DeQueue() (int, bool) {
	if q.size == 0 {
		return -1, false
	}
	data := q.head.next.data
	q.head.next = q.head.next.next
	if q.head.next == nil {
		q.rear = q.head
	}
	q.size--
	return data, true
}

//队头
func (q linkQueue) GetHead() (int, bool) {
	if q.size == 0 {
		return -1, false
	}
	return q.head.next.data, true
}
