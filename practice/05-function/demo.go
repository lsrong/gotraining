package main

import "fmt"

/*
func main() {
	// 1-100的质数
	demoPrime()

	// 100-1000的水仙数
	demoNarcissisticNumber()

	// 统计一段字符串的英文字母，数字，空格，其他字符的数量
	enC, numC, spaceC, otherC := demoStatChar("520 1314  demo 示例")
	fmt.Printf("英文字符：%d \n", enC)
	fmt.Printf("数字字符：%d \n", numC)
	fmt.Printf("空格字符：%d \n", spaceC)
	fmt.Printf("其他字符：%d \n", otherC)
}
*/

// demoPrime 求1到100之内的所有质数
func demoPrime() {
	for i := 1; i <= 100; i++ {
		if isPrime(i) {
			fmt.Printf("i = %d is prime \n", i)
		}
	}
}

// isPrime 判断一个数字是否为质数
func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	// 只能1和本身整除的数值
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

// demoNarcissisticNumber 求[100, 1000)之间的水仙花数
// 水仙花数：每个位上面的数字的3次幂之和等于它本身，例如：1^3+5^3+3^3=153
func demoNarcissisticNumber() {
	for number := 100; number < 1000; number++ {
		if isNarcissisticNumber(number) {
			fmt.Printf("number=%d is a narcissistic number \n", number)
		}
	}
}

// isNarcissisticNumber 判断一个数是否为水仙花数
func isNarcissisticNumber(number int) bool {
	f, s, t := number%10, (number/10)%10, (number/100)%10
	sum := f*f*f + s*s*s + t*t*t

	return sum == number
}

// statChar 统计一段字符中的英文字母，数字，空格，其他字符的数量
func demoStatChar(str string) (int, int, int, int) {
	var enC, numC, spaceC, otherC int
	// 字符串转成 []rune
	chars := []rune(str)
	fmt.Printf("统计字符为 %s \n", string(chars))
	for i := 0; i < len(chars); i++ {
		temp := chars[i]
		switch {
		case temp >= 'a' && temp <= 'z' || temp >= 'A' && temp <= 'Z':
			enC++
		case temp > '0' && temp <= '9':
			numC++
		case temp == ' ':
			spaceC++
		default:
			otherC++
		}
	}
	return enC, numC, spaceC, otherC
}
