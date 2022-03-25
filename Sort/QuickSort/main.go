package main

import "fmt"

//快速排序，先把第一个元素当成基准，把比他大的放在右边，比他小的放在左边
type uint64Slice []uint64

func main() {
	numbers := []uint64{5, 4, 20, 3, 8, 2, 8}
	quickSort(numbers, 0, len(numbers)-1)
	fmt.Println(numbers)
}

func quickSort(numbers uint64Slice, start, end int) {
	var middle int
	tempStart := start
	tempEnd := end

	if tempStart >= tempEnd {
		return
	}
	pivot := numbers[start]
	for start < end {
		//查找在右边比他小的元素
		for start < end && numbers[end] > pivot {
			end--
		}
		if start < end {
			numbers.swap(start, end)
			start++
		}
		//查找左边比他大的元素
		for start < end && numbers[start] < pivot {
			start++
		}
		if start < end {
			numbers.swap(start, end)
			end--
		}
	}
	//将基准放入最中间
	numbers[start] = pivot
	//记录中间值的下标，为后面的快速排序做准备
	middle = start

	quickSort(numbers, tempStart, middle-1)
	quickSort(numbers, middle+1, tempEnd)

}

// 交换方法
func (numbers uint64Slice) swap(i, j int) {
	numbers[i], numbers[j] = numbers[j], numbers[i]
}
