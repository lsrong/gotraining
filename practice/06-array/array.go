package main

import "fmt"

func main() {
	demoArray()
	demoIterate()
	demoMultiArrays()

	b := [5]int{1, 2, 3, 4, 5}
	sum := practiceSumIntArray(b)
	fmt.Printf("practiceSumIntArray sum = %d \n", sum)

	c := [5]int{1, 3, 5, 8, 7}
	findTwoSumElement(c, 10)
}

// demoArray 定义数组类型： 长度 + 类型， 从下标0开始
// 注意：长度和类型 一起组成go语言中的数组类型
func demoArray() {
	var a [10]int
	a[0], a[1] = 1, 2
	fmt.Printf("base init a = %v \n", a)

	// 定义时候赋值
	var b = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("base init b = %v \n", b)
	c := [5]int{6, 7, 8, 9, 10}
	fmt.Printf("base init c = %v \n", c)

	// [...] :长度为定义的的最长下标数 + 1
	d := [...]int{0, 1, 2, 3, 4} //  index = 5
	fmt.Printf("base init d = %v \n", d)

	// 初始化指定下标， index:value, 其他未初始化下标默认为零值
	e := [...]int{4: 10, 6: 20}
	fmt.Printf("base init e = %v \n", e)
	fmt.Printf("base init len(e) = %d \n", len(e))
}

// demoIterate 数组遍历
// 1. for 2.for range
func demoIterate() {
	nums := [5]int{0, 1, 2, 3, 4}
	// for i:=0; i< len(nums); i++ {}
	for i := 0; i < len(nums); i++ {
		fmt.Printf("for nums[%d] = %d \n", i, nums[i])
	}
	// for i,v :=range {}
	for i, v := range nums {
		fmt.Printf("for range nums[%d] = %d \n", i, v)
	}

	// for _ v :=range {}
	for _, v := range nums {
		fmt.Printf("for  _ v = %d \n", v)
	}
}

// demoMultiArrays 多维数组， 定义 + 遍历
func demoMultiArrays() {
	// 多维数组的定义
	nums := [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	fmt.Printf("mutli array: %v \n", nums)

	// for 遍历
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			fmt.Printf("mutli array for nums[%d][%d] = %d \n", i, j, nums[i][j])
		}
	}

	// for range 遍历
	for i, v1 := range nums {
		for j, v2 := range v1 {
			fmt.Printf("mutli array for range nums[%d][%d] = %d \n", i, j, v2)
		}
	}
}

// practiceSumIntArray 求数组所有元素的和
func practiceSumIntArray(a [5]int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}

// findTwoSumElement 找出数组中和为给定值的两个元素的下标，⽐如数组:[1,3,5,8,7]，找出两个元素之和等于8的下标分别是(0, 4)和(1,2)。
func findTwoSumElement(a [5]int, total int) {
	for i, v1 := range a {
		target := total - v1
		for j := i + 1; j < len(a); j++ {
			if target == a[j] {
				fmt.Printf("findTwoSumElement (%d, %d)\n", i, j)
			}
		}
	}
}
