package main

import "fmt"

/* 求1到100之内的所有质数 */
func getPrime() {
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Printf("i=%d is prime \n", i)
		}
	}
}

// 判断是否为质数
func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

/* 求100,1000，之间的水仙花数 */
// 水仙花数是指一个 3 位数，它的每个位上的数字的 3次幂之和等于它本身（例如：1^3 + 5^3+ 3^3 = 153）
func getNarcissisticNumber() {
	for number := 100; number < 1000; number++ {
		if isNarcissisticNumber(number) {
			fmt.Printf("number=%d is narcissistic number \n", number)
		}
	}
}

// 判断是否为水仙数
func isNarcissisticNumber(number int) bool {
	first, second, third := number%10, (number/10)%10, (number/100)%10
	//fmt.Println(first, second, third)
	sum := first*first*first + second*second*second + third*third*third
	return sum == number
}

/* 统计一段字符中的英文字母，数字，空格，其他字符的数量 */
func statChar(str string) (enChar int, numChar int, spaceChar int, otherChar int) {
	chars := []rune(str)
	fmt.Println(chars)
	for i := 0; i < len(chars); i++ {
		// 英文
		if chars[i] >= 'a' && chars[i] <= 'z' || chars[i] >= 'A' && chars[i] <= 'Z' {
			enChar++
			continue
		}
		// 数字
		if chars[i] >= '0' && chars[i] <= '9' {
			numChar++
			continue
		}

		// 空格
		if chars[i] == ' ' {
			spaceChar++
			continue
		}

		otherChar++
	}
	return
}
func main() {
	// 求1-100质数
	getPrime()

	// 求100-1000的水仙数
	getNarcissisticNumber()

	// 统计一段字符中的英文字母，数字，空格，其他字符的数量
	str := "abcd   测试 123"
	enChar, numChar, spaceChar, otherChar := statChar(str)
	fmt.Printf("英文字符：%d \n", enChar)
	fmt.Printf("数字字符：%d \n", numChar)
	fmt.Printf("空格字符：%d \n", spaceChar)
	fmt.Printf("其他字符：%d \n", otherChar)
}
