package sort

func SelectSort(num []int) {
	for i := 0; i < len(num)-1; i++ {
		index := i
		for j := i + 1; j < len(num); j++ {
			if num[index] > num[j] {
				index = j
			}
		}
		temp := num[i]
		num[i] = num[index]
		num[index] = temp
	}
}
