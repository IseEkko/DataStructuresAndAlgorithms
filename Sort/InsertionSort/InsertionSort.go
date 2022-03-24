package main

import "fmt"

func main() {
	res := make([]int, 0)
	res = []int{
		2, 1, 4, 5, 3,
	}
	InsertionSort(res, len(res))
	fmt.Println(res)
}

//插入排序
func InsertionSort(nums []int, n int) {
	if n < 1 {
		return
	}
	for i := 0; i < n; i++ {
		value := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = value
	}
}
