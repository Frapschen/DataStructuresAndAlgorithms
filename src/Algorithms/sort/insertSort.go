package sort

func InsertSort(num []int) {
	for i := 1; i < len(num); i++ {
		j := i - 1
		temp := num[i]
		for ; j >= 0 && num[j] > temp; j-- {
			num[j+1] = num[j]
		}
		num[j+1] = temp
	}
}
