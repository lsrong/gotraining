package main

import "fmt"

// Sample program to show the components of a slice. It has a
// length, capacity and the underlying array.

// 切片的组成: 长度, 容量, 底层数组, 切片是底层数据的引用,因此切片传值也是引用类型.
/* type slice struct {
	ptr *Elem
	len int
	cap int
}
type Elem [cap]T
*/

func main() {
	fruits := make([]string, 5, 8)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	for i, s := range fruits {
		fmt.Printf("Outside Func: Index[%d] Address[%p] Vaule[%s]\n", i, &fruits[i], s)
	}

	// pass slice by pointer
	inspectSlice(fruits)

}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d] \n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("Inside Func: Isndex[%d] Address[%p] Vaule[%s]\n", i, &slice[i], s)
	}
}
