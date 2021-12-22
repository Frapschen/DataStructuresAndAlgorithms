package search

//查找搜索 需要数据有序
func InterpolationSearch(data []int, number int) int {
	if len(data) == 0 {
		return -1
	}
	mid := 0
	low := 0
	high := len(data) - 1
	for low <= high {
		mid = low + ((number-data[low])/(high-low))*(high-low)
		if data[mid] == number {
			return mid
		} else if data[mid] > number {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return -1
}
