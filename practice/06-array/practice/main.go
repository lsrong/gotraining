package main

import "fmt"

// 求数组所有元素之和
func sumIntArray(a [5]int) (sum int) {
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return
}

// 找出数组中和为给定值的两个元素的下标，⽐比如数组:[1,3,5,8,7]，找出两个元素之和等于8的下标分别是(0, 4)和(1,2)。
func findTwoSumElement(a [5]int, total int) {
	l := len(a)
	for i := 0; i < l; i++ {
		target := total - a[i]
		for j := i + 1; j < l; j++ {
			if target == a[j] {
				fmt.Printf("(%d, %d) \n", i, j)
			}
		}
	}
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sumIntArray(a))

	b := [5]int{1, 3, 5, 8, 7}
	findTwoSumElement(b, 8)
}
