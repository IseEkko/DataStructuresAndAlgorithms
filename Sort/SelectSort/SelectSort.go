package main

import "fmt"

type uint64Slice []uint64

func main() {
	numbers := []uint64{5, 23, 1, 6, 7, 9, 2}
	selectSort(numbers)
	fmt.Println(numbers)
}

//选择排序，选择数组中最小的部分，然后放在最前面，然后在剩下的里面选择最小的再放在最前面
func selectSort(numbers uint64Slice) {
	for i := 0; i < len(numbers)-1; i++ {
		// 记录最小值位置
		min := i
		for j := i + 1; j < len(numbers); j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}
		if i != min {
			numbers.swap(i, min)
		}
	}
}

// 交换方法
func (numbers uint64Slice) swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
