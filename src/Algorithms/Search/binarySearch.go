package search

//需要数据有序
func binarySearch(data []int, number int) int {
	if len(data) == 0 {
		return -1
	}
	mid := 0
	low := 0
	high := len(data) - 1
	for low <= high {
		mid = (high + low) / 2
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
