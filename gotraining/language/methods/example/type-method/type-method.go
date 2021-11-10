package main

import "fmt"

// Sample program to show how to declare methods against a named type.
// 演示如何针对命名类型声明方法的示例程序。
// 为命名类型声明方法

// "If the method needs to mutate the receiver, the receiver must be a pointer."
//  go wiki: https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
//  如果方法需要改变接收者，接收者必须是一个指针.

type duration int64

const (
	nanosecond  duration = 1
	microsecond          = 1000 * nanosecond
	millisecond          = 1000 * microsecond
	second               = 1000 * millisecond
	minute               = 60 * second
	hour                 = 60 * minute
)

func (d *duration) serHours(h float64) {
	*d = duration(h) * hour
}

func (d duration) hours() float64 {
	h := d / hour
	nsec := d % hour

	return float64(h) + float64(nsec)*(1e-9/60/60)
}

func main() {
	var dur duration

	dur.serHours(8)

	fmt.Println("Hours: ", dur.hours())
}
