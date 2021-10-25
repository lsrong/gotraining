package main

import "fmt"

// switch var {cast condition:}
func testSwitch() {
	a := 4
	switch a {
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	case 3:
		fmt.Println("a=3")
	case 4:
		fmt.Println("a=4")
	case 5:
		fmt.Println("a=5")
	}
}

func getValue() int {
	return 10
}

// switch assignment;var {case condition}
func testSwitchAssignment() {
	switch a := getValue(); a {
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Println("a=2")
	case 3:
		fmt.Println("a=3")
	case 4:
		fmt.Println("a=4")
	case 5:
		fmt.Println("a=5")
	default:
		fmt.Println("a is plus 5")
	}

}

// case 多个条件
func testCaseMulti() {
	switch a := getValue(); a {
	case 1, 2, 3, 4, 5:
		fmt.Println("a >=1 and a <= 5")
	case 6, 7, 8, 9, 10:
		fmt.Println("a >=6 and a <= 10")
	default:
		fmt.Println("a <=0 or a > 10")
	}
}

// case 范围
func testRangeSwitch(num int) {
	switch {
	case num >= 0 && num <= 25:
		fmt.Println("a >= 0 and a <= 25")
	case num > 25 && num <= 50:
		fmt.Println("a > 25 and a <= 50")
	case num > 50 && num <= 75:
		fmt.Println("a > 50 and a <= 75")
	case num > 75 && num <= 100:
		fmt.Println("a > 75 and a <= 100")
	default:
		fmt.Println("The num is not in (0,100)")
	}
}

func main() {
	testSwitch()

	testSwitchAssignment()

	testCaseMulti()

	testRangeSwitch(77)
}
