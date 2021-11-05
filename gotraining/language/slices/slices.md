## Slices - Arrays, Slices and Maps

Slices are an incredibly important data structure in Go. They form the basis for how we manage and manipulate data in a flexible, performant and dynamic way. It is incredibly important for all Go programmers to learn how to uses slices.

切片是 Go 中非常重要的数据结构。它们构成了我们如何以灵活、高效和动态的方式管理和操作数据的基础。对于所有 Go 程序员来说，学习如何使用切片非常重要。

## Notes
* A slice is not an array. A slice describes a piece of an array
* 切片不是数组。切片描述了数组的一部分

* Slices are like dynamic arrays with special and built-in functionality.  
* 切片就像具有特殊和内置功能的动态数组。

* There is a difference between a slices length and capacity and they each service a purpose.
* 这是切片长度和容量之间的差异，它们各自服务于一个目的。

* Slices allow for multiple "views" of the same underlying array.
* 切片允许同一底层数组的多个“视图” "slice[n:m]"。

* Slices can grow through the use of the built-in function append.
* 切片可以通过使用内置函数 append 增长。

## Links

[Go Slices: usage and internals](https://blog.golang.org/go-slices-usage-and-internals) - Andrew Gerrand    
[Strings, bytes, runes and characters in Go](https://blog.golang.org/strings) - Rob Pike    
[Arrays, slices (and strings): The mechanics of 'append'](https://blog.golang.org/slices) - Rob Pike    
[Understanding Slices in Go Programming](https://www.ardanlabs.com/blog/2013/08/understanding-slices-in-go-programming.html) - William Kennedy    
[Collections Of Unknown Length in Go](https://www.ardanlabs.com/blog/2013/08/collections-of-unknown-length-in-go.html) - William Kennedy    
[Iterating Over Slices In Go](https://www.ardanlabs.com/blog/2013/09/iterating-over-slices-in-go.html) - William Kennedy    
[Slices of Slices of Slices in Go](https://www.ardanlabs.com/blog/2013/09/slices-of-slices-of-slices-in-go.html) - William Kennedy    
[Three-Index Slices in Go 1.2](https://www.ardanlabs.com/blog/2013/12/three-index-slices-in-go-12.html) - William Kennedy    
[SliceTricks](https://github.com/golang/go/wiki/SliceTricks)

## Code Review

[Declare and Length](example/declaration/declaration.go)  
[Reference Types](example/reference_type/reference_type.go)  
[Appending slices](example/append/append.go)  
[Taking slices of slices](example/slice_of_slice/sliceOfSlice.go)  
[Slices and References](example/share_reference/share_reference.go)  
[Strings and slices](example/string_slice/string_slice.go)  
[Variadic functions](example/variadic_function/variadic_function.go)  
[Range mechanics](example/range_mechanics/range_mechanics.go)  
["Gotcha"](example/gotcha/gotcha.go)  

## Advanced Code Review

[Three index slicing](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/2CM_LPBnfIR))

## Exercises

### Exercise 1

**Part A** Declare a nil slice of integers. Create a loop that appends 10 values to the slice. Iterate over the slice and display each value.
// 声明一个 nil 整数切片。创建一个循环，将 10 个值附加到切片。迭代切片并显示每个值

**Part B** Declare a slice of five strings and initialize the slice with string literal values. Display all the elements. Take a slice of index one and two and display the index position and value of each element in the new slice.
// 声明一个包含五个字符串的切片，并使用字符串文字值初始化切片。显示所有元素。取索引一、二的切片，显示新切片中每个元素的索引位置和值。
[Exercise](exercise/exercise.go)
