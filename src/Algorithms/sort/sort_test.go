package sort

import (
	"fmt"
	"testing"
)

func getNum() []int {
	num := []int{1, 7, 3, 9, 5}
	return num
}

func Test_BubbleSort(t *testing.T) {
	num := getNum()
	BubbleSort(num)
	fmt.Println(num)
}

func Test_insertSort(t *testing.T) {
	num := getNum()
	InsertSort(num)
	fmt.Println(num)
}

func Test_SelectSort(t *testing.T) {
	num := getNum()
	SelectSort(num)
	fmt.Println(num)
}

func Test_quickSort(t *testing.T) {
	num := getNum()
	QuickSort(num, 0, len(num)-1)
	fmt.Println(num)
}

func Test_shellSort(t *testing.T) {
	num := getNum()
	ShellSort(num)
	fmt.Println(num)
}

func Test_heapSort(t *testing.T) {
	num := getNum()
	HeapSort(num)
	fmt.Println(num)
}

func Test_MergeSort(t *testing.T) {
	num := getNum()
	temp := make([]int, len(num))
	MergeSort(num, 0, len(num)-1, temp)
	fmt.Println(num)
}
