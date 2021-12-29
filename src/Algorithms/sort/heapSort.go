package sort

func HeapSort(num []int) {
	//初始化堆(构建大顶堆)
	//最后一个有孩子的节点：i := (len(num)-1)/2,从这向上依次调整为大顶堆
	for i := (len(num) - 1) / 2; i >= 0; i-- {
		heapAdjustB(num, i, len(num))
	}
	//交换堆顶与堆尾元素
	for i := len(num) - 1; i > 0; i-- {
		temp := num[i]
		num[i] = num[0]
		num[0] = temp
		//再次调整
		heapAdjustB(num, 0, i)
	}
}

func heapAdjustB(num []int, s int, len int) {
	temp := num[s]
	child := s*2 + 1
	for child < len {
		//child 为右孩子，如果右孩子大，那就选它作为交换对象
		if child+1 < len && num[child+1] > num[child] {
			child++
		}
		if num[child] > temp {
			num[s] = num[child]
			num[child] = temp
			//待调整结点移动到较大子孩子原来的位置
			//并再找其孩子节点，如果有继续调整
			s = child
			child = s*2 + 1
		} else {
			//待调整节点为最大值
			break
		}
	}
}
