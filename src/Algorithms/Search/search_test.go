package search

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	//有序
	data := []int{1, 2, 3, 5, 6, 7, 8, 9, 10, 13, 15}
	// fmt.Println("线性：", LinearSearch(data, 7))
	// fmt.Println("二分：", binarySearch(data, 7))
	// fmt.Println("插值：", InterpolationSearch(data, 7))
	fmt.Println("指数：", ExponentialSearch(data, 15))

}
