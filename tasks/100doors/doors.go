package main

import (
	"fmt"
)

type data struct {
	door []bool
}

type Toggler interface {
	toggle(d *data)
}

// Unoptimized 未优化的实现, 按照任务的正常思路, 两层循环: 第一层 执行 n 次, 第二层, 执行 n/i 次
type Unoptimized struct{}

func (u *Unoptimized) toggle(d *data) {
	number := len(d.door)
	for i := 1; i <= number; i++ {
		for j := i - 1; j < number; j += i {
			d.door[j] = !d.door[j]
		}
	}
}

// Optimized1 优化1, 从输出的结果中总结出的规律, 开启的门 = prev + (2*n +1), 注: prev为上一次的数值, n 为开门计数
type Optimized1 struct{}

func (o *Optimized1) toggle(d *data) {
	door := 1
	incrementer := 0
	for current := 1; current <= len(d.door); current++ {
		if current == door {
			d.door[current-1] = true
			incrementer++
			door += (2 * incrementer) + 1
		}
	}
}

// Optimized2 优化2, 每次打开门的数值是n的平方, 整数的完美平方, n^2
type Optimized2 struct{}

func (o *Optimized2) toggle(d *data) {
	/*
		for i := 1; i <= len(d.door); i++ {
			sqrt := math.Sqrt(float64(i)) // Math.Sqrt(): 求平方根
			if math.Mod(sqrt, 1) == 0 {   // Math.Mod():  取模
				d.door[i-1] = true
			}
		}
	*/
	incrementer := 1
	for i := 1; i <= len(d.door); i++ {
		if incrementer*incrementer == i {
			incrementer++
			d.door[i-1] = true
		}
	}
}

func main() {
	d := &data{make([]bool, 100)}
	ts := []Toggler{
		new(Unoptimized),
		new(Optimized1),
		new(Optimized2),
	}

	progress(ts, d)
}

func progress(ts []Toggler, d *data) {
	for _, t := range ts {
		reset(d)
		t.toggle(d)
		printDoors(d.door)
		fmt.Println("------------ ")
	}
}

func reset(d *data) {
	for i := range d.door {
		d.door[i] = false
	}
}

func printDoors(dr []bool) {
	for i, v := range dr {
		if v {
			fmt.Printf("(%d)", i+1)
		} else {
			fmt.Print(0)
		}
		if i%10 == 9 {
			fmt.Print("\n")
		} else {
			fmt.Print(" ")
		}
	}
}
