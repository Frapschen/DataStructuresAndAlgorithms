package sort

func ShellSort(num []int) {
	for gap := len(num) / 2; gap > 0; gap /= 2 {
		for i := gap; i < len(num); i++ {
			j := i - gap
			temp := num[i]
			for ; j >= 0 && num[j] > temp; j -= gap {
				num[j+gap] = num[j]
			}
			num[j+gap] = temp
		}
	}
}
