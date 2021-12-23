package caching

import "fmt"

// 演示代码来说明为什么面向数据的设计很重要。数据布局对性能的影响比对算法效率的影响更大。

// 分别使用链表和二维数组保存一组大数据，测试跳跃访问偶数元素的效率

const (
	cols = 4 * 1024
	rows = 4 * 1024
)

// data 链表元素结构体
type data struct {
	v    byte
	next *data
}

// matrix 代表二维大矩阵数组变量
var matrix [rows][cols]byte

var list *data

func init() {
	last := list
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			var d data
			if list == nil {
				list = &d
			}
			if last != nil {
				last.next = &d
			}
			last = &d

			// 为偶数行时候标记为0xFF
			if row%2 == 0 {
				matrix[row][col] = 0xFF
				d.v = 0xFF
			}
		}
	}

	var counter int
	d := list
	for d != nil {
		counter++
		d = d.next
	}
	fmt.Printf("Elements in link list: %d \n", counter)
	fmt.Printf("Elements in  matrix: %d \n", rows*cols)
}

// LinkedListTraverse 线性遍历链表。
func LinkedListTraverse() int {
	var counter int // 记录0xFF数量
	d := list
	for d != nil {
		if d.v == 0xFF {
			counter++
		}
		d = d.next
	}

	return counter
}

// ColumnTraverse 沿每一列线性遍历矩阵
func ColumnTraverse() int {
	var counter int
	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if matrix[row][col] == 0xFF {
				counter++
			}
		}
	}
	return counter
}

// RowTraverse 沿每一行线性遍历矩阵
func RowTraverse() int {
	var counter int
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if matrix[row][col] == 0xFF {
				counter++
			}
		}
	}

	return counter
}
