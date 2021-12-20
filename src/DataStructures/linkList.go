package datastructures

type linkNode struct {
	next *linkNode
	data int
}

//带头节点的单链表
type linkList struct {
	size int
	head *linkNode
}

//初始化
func Init() linkList {
	return linkList{size: 0, head: &linkNode{}}
}

//增加一个节点,使用头插法
func (l *linkList) AddNode(data int) {
	newNode := &linkNode{data: data,
		next: l.head.next}
	l.head.next = newNode
	l.size = l.size + 1
	return
}

//删除指定x的所有节点
func (l *linkList) DelNode(data int) {
	begin := l.head.next
	pre := l.head
	for begin != nil {
		if begin.data == data {
			pre.next = begin.next
			begin = begin.next
			l.size = l.size - 1
		} else {
			pre = begin
			begin = begin.next
		}
	}
	// for begin != nil {
	// 	if begin.data == data {
	// 		temp := begin.next
	// 		if temp != nil {
	// 			begin.data = temp.data
	// 			begin.next = temp.next
	// 			temp = nil
	// 		} else {
	// 			begin = nil //因为是值拷贝 无法让原来的指针变空，也没有主动释放内存的函数
	// 		}
	// 		l.size = l.size - 1
	// 	}
	// 	if begin == nil {
	// 		return
	// 	}
	// 	begin = begin.next
	// }
	// return

	// for begin := l.head.next; begin != nil; begin = begin.next {
	// 	if begin.data == data {
	// 		temp := begin.next
	// 		if temp != nil {
	// 			begin.data = temp.data
	// 			begin.next = temp.next
	// 			temp = nil
	// 		} else {
	// 			begin = nil
	// 			l.size = l.size - 1
	// 			return
	// 		}
	// 		l.size = l.size - 1
	// 	}
	// }
	// return
}

//查询一个节点
func (l linkList) Find(data int) bool {
	begin := l.head.next
	for begin != nil {
		if begin.data == data {
			return true
		}
		begin = begin.next
	}
	return false
}

//遍历所有节点
func (l linkList) ToSlice() []int {
	var dataSlice []int
	for begin := l.head.next; begin != nil; begin = begin.next {
		dataSlice = append(dataSlice, begin.data)
	}
	return dataSlice
}
