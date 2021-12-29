package sort

func MergeSort(num []int, low int, high int, temp []int) {
	if low < high {
		mid := (high + low) / 2
		MergeSort(num, low, mid, temp)
		MergeSort(num, mid+1, high, temp)
		merge(num, low, mid, high, temp)
	}

}

//比较low和high的大小，放入temp中，完成一次归并
func merge(num []int, low int, mid int, high int, temp []int) {
	i := low
	j := mid + 1
	k := 0
	for i <= mid && j <= high {
		if num[i] > num[j] {
			temp[k] = num[j]
			k++
			j++
		} else {
			temp[k] = num[i]
			k++
			i++
		}
	}
	for i <= mid {
		temp[k] = num[i]
		k++
		i++
	}
	for j <= high {
		temp[k] = num[j]
		k++
		j++
	}
	k = 0
	for low <= high {
		num[low] = temp[k]
		low++
		k++
	}
}
