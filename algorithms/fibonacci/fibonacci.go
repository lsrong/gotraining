package fibonacci

/**
斐波纳契数（Fibonacci sequence），又称黄金分割数列，因数学家列昂纳多·斐波那契（Leonardoda Fibonacci）以兔子繁殖为例子而引入，故又称为“兔子数列”，
指的是这样一个数列：0、1、1、2、3、5、8、13、21、34、……

用文字来说，就是斐波那契数列由0和1开始，之后的斐波那契数就是由之前的两数相加而得出。
在数学上，斐波纳契数列以如下被以递归的方法定义：F(0)=0，F(1)=1, F(n)=F(n-1)+F(n-2)（n>=2，n∈N*）
*/

type FibBuilder func(n int) int

func BuildFib(n int, f FibBuilder) []int {
	var fib []int
	for i := 0; i < n; i++ {
		fib = append(fib, f(i))
	}
	return fib
}

// NewFibViaRecursion 递归实现
func NewFibViaRecursion(n int) []int {
	return BuildFib(n, fibRecursion)
}

// fibRecursion 时间复杂度为[O(2^n)]。
func fibRecursion(n int) int {
	if n < 2 {
		return n
	}
	return fibRecursion(n-1) + fibRecursion(n-2)
}

// NewFibViaDp 动态规划实现。
func NewFibViaDp(n int) []int {
	return BuildFib(n, fibDp)
}

// fibDp 时间复杂度为O(n）
func fibDp(n int) int {
	dp := map[int]int{0: 0, 1: 1}
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
