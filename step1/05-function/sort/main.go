package main

import "fmt"

// 使用快速排序、选择排序、冒泡排序算法对数字数组进行升序排序

/* 快速排序 */
func insertSort(a [8]int) [8]int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
	return a
}

/* 选择排序 */
func selectSort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

/* 冒泡排序 */
func bubbleSort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	a := [8]int{10, 1, 8, 9, 4, 6, 7, 2}
	fmt.Println(insertSort(a))
	fmt.Println(selectSort(a))
	fmt.Println(bubbleSort(a))
}
