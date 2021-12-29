package sort

func QuickSort(num []int, low int, high int) {
	if low >= high {
		return
	}
	l := low
	h := high
	mid := num[l]
	//每一趟都分割出两个子列出来
	for l < h {
		for l < h && num[h] >= mid {
			h--
		}
		num[l] = num[h]
		for l < h && num[l] <= mid {
			l++
		}
		num[h] = num[l]
	}
	num[l] = mid
	QuickSort(num, 0, l-1)
	QuickSort(num, l+1, high)
	return
}
