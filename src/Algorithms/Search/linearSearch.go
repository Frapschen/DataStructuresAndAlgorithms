package search

func LinearSearch(data []int, number int) int {
	if len(data) == 0 {
		return -1
	}
	for i, value := range data {
		if value == number {
			return i
		}
	}
	return -1
}
