package datastructures

const MAXSIZE int = 5

type cQueue struct {
	queue       [MAXSIZE]int
	front, rear int
}

//初始化
func CQueueInit() cQueue {
	return cQueue{}
}

//入队
func (c *cQueue) CQueueEnQueue(data int) bool {
	if (c.rear+1)%MAXSIZE == c.front {
		return false
	}
	c.queue[c.rear] = data
	c.rear = (c.rear + 1) % MAXSIZE
	return true
}

//出队
func (c *cQueue) CQueueDeQueue() (int, bool) {
	if c.rear == c.front {
		return -1, false
	}
	data := c.queue[c.front]
	c.front = (c.front + 1) % MAXSIZE
	return data, true
}

//队头
func (c cQueue) CQueueGetHead() (int, bool) {
	if c.rear == c.front {
		return -1, false
	}
	return c.queue[c.front], true
}
