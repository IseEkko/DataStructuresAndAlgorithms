package main

import (
	"fmt"
	"math"
)

type uint64Slice []uint64

func main() {
	numbers := []uint64{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}
	shellSort(numbers)
	fmt.Println(numbers)
}

func shellSort(numbers uint64Slice) {
	gap := 1
	for gap < len(numbers) {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < len(numbers); i++ {
			//记录数字方便后面的交换
			tmp := numbers[i]
			//确定前一个数字的下标
			j := i - gap
			for j >= 0 && numbers[j] > tmp {
				numbers[j+gap] = numbers[j]
				j -= gap
			}
			//交换数据
			numbers[j+gap] = tmp
		}
		gap = int(math.Floor(float64(gap / 3)))
	}
}
