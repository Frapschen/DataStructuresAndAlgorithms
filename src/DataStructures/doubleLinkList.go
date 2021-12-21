package datastructures

type dNode struct {
	data        int
	next, prior *dNode
}
type doubleLinkList struct {
	size int
	head *dNode
}

//初始化
func DLinkListInit() doubleLinkList {
	return doubleLinkList{size: 0,
		head: &dNode{}}
}

//插入
func (d *doubleLinkList) AddDNode(data int) {
	newNode := &dNode{data: data,
		prior: d.head,
		next:  d.head.next}
	if d.head.next != nil {
		d.head.next.prior = newNode
	}
	d.head.next = newNode
	d.size++
}

//遍历 返回两个参数，一个是正向，一个是逆向
func (d doubleLinkList) DLinkListToSlice() ([]int, []int) {
	begin := d.head.next
	var data []int
	var data2 []int
	for begin != nil {
		data = append(data, begin.data)
		begin = begin.next
	}
	begin = d.head.next
	for begin.next != nil {
		begin = begin.next
	}
	for begin != d.head {
		data2 = append(data2, begin.data)
		begin = begin.prior
	}
	return data, data2
}

//删除指定节点
func (d *doubleLinkList) DelDLinkList(data int) {
	begin := d.head.next
	for begin != nil {
		if begin.data == data {
			if begin.next == nil {
				begin.prior.next = nil
				return
			}
			begin.prior.next = begin.next
			begin.next.prior = begin.prior
			d.size--
		}
		begin = begin.next
	}
}

//查找指定节点
func (d doubleLinkList) FindDLinkList(data int) bool {
	for begin := d.head.next; begin != nil; begin = begin.next {
		if begin.data == data {
			return true
		}
	}
	return false
}
