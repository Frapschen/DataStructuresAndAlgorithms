package search

import (
	"math"
)

//指数搜索 需要数据有序
func ExponentialSearch(data []int, number int) int {
	maxLenth := len(data)
	switch maxLenth {
	case 0:
		return -1
	case 1:
		if data[0] == number {
			return 0
		}
	case 2:
		if data[0] == number {
			return 0

		}
		if data[1] == number {
			return 1
		}
	default:
		if data[0] == number {
			return 0

		}
		if data[1] == number {
			return 1
		}
	}
	i := 1
	index := 0
	for {
		index = int(math.Exp2(float64(i)))
		if index >= maxLenth-1 {
			return search(data, number, int(math.Exp2(float64(i-1))), maxLenth-1)
		} else if data[index] == number {
			return index
		} else if data[index] > number {
			return search(data, number, int(math.Exp2(float64(i-1))), index)
		}
		i++
	}
}

func search(data []int, number int, low int, high int) int {
	if len(data) == 0 {
		return -1
	}
	mid := 0
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
