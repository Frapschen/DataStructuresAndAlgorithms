package offer

/*
题目一：找出数组中重复的数字
*/
func func1(n int, data ...int) int {
	for i := 0; i < n; i++ {
		if data[i] != i {
			temp := data[i]
			data[i] = data[data[i]]
			data[data[i]] = temp
		}
	}
}
