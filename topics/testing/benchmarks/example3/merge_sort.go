package example3

// 用基准测试测试归并排序的不同实现之间的性能差异.

import (
	"math"
	"runtime"
	"sync"
)

// Single 使用单个 goroutine 来执行归并排序.goroutine: 1
func Single(n []int) []int {
	if len(n) <= 1 {
		return n
	}

	i := len(n) / 2

	l := Single(n[:i])

	r := Single(n[i:])

	return merge(l, r)
}

// Unlimited 对每个拆分使用一个 goroutine 来执行合并排序。goroutine: 1, 2, 4, 8, 32, 64 ....
func Unlimited(n []int) []int {
	if len(n) <= 1 {
		return n
	}
	i := len(n) / 2
	var l, r []int
	var wg sync.WaitGroup
	wg.Add(2)
	// 并发排序左边数值
	go func() {
		l = Unlimited(n[:i])
		wg.Done()
	}()

	// 并发排序右边数值
	go func() {
		r = Unlimited(n[i:])
		wg.Done()
	}()

	wg.Wait()

	return merge(l, r)
}

// NumCPU 使用与我们拥有的内核相同数量的 goroutines 来执行合并排序。1, 2, 4, 8, 1, 1 ...
func NumCPU(n []int, lvl int) []int {
	if len(n) <= 1 {
		return n
	}
	i := len(n) / 2
	var l, r []int

	// Calculate how many levels deep we can create goroutines.
	// On an 8 core machine we can keep creating goroutines until level 4.
	// 		Lvl 0		1  Lists		1  Goroutine
	//		Lvl 1		2  Lists		2  Goroutines
	//		Lvl 2		4  Lists		4  Goroutines
	//		Lvl 3		8  Lists		8  Goroutines
	//		Lvl 4		16 Lists		16 Goroutines
	maxLevel := int(math.Log2(float64(runtime.NumCPU())))
	for lvl <= maxLevel {
		lvl++
		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			l = NumCPU(n[:i], lvl)
			wg.Done()
		}()

		go func() {
			r = NumCPU(n[i:], lvl)
			wg.Done()
		}()

		wg.Wait()

		return merge(l, r)
	}

	l = NumCPU(n[:i], lvl)
	r = NumCPU(n[i:], lvl)

	return merge(l, r)
}

func merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for {
		switch {
		case len(l) == 0:
			// 左边为空，批量添加右边的元素
			return append(ret, r...)

		case len(r) == 0:
			// 右边为空，批量添加左边的元素
			return append(ret, l...)

		case l[0] <= r[0]:
			ret = append(ret, l[0])

			l = l[1:]
		default:
			ret = append(ret, r[0])

			r = r[1:]
		}
	}
}
