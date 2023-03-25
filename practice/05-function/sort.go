package main

import "fmt"

func main() {
	nums := [8]int{10, 1, 8, 9, 4, 6, 7, 2}
	fmt.Printf("冒泡排序： %v \n", demoBubbleSort(nums))
	fmt.Printf("选择排序： %v \n", demoSelectSort(nums))
	fmt.Printf("快速排序： %v \n", insertSort(nums))
}

// demoBubbleSort 冒泡排序
func demoBubbleSort(nums [8]int) [8]int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// demoSelectSort 选择排序
func demoSelectSort(nums [8]int) [8]int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return nums
}

// 快速排序
func insertSort(nums [8]int) [8]int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			} else {
				break
			}
		}
	}
	return nums
}
