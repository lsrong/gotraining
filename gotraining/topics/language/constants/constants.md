## Constants

Constants are a way to create a named identifier whose value can never change. They also provide an incredible amount of flexibility to the language. The way constants are implemented in Go is very unique.  
常量是一种创建命名标识符的方法，其值永远不会改变。它们还为语言提供了难以置信的灵活性。 Go 中常量的实现方式非常独特。

## Notes

* Constants are not variables.
* 常量不是变量

* They exist only at compilation.
* 它们仅在编译时存在。

* Untyped constants can be implicitly converted where typed constants and variables can't.
* 无类型常量可以隐式转换，而类型常量和变量不能。

* Think of untyped constants as having a Kind, not a Type.
* 将无类型常量视为具有种类，而不是类型。

* Learn about explicit and implicit conversions.
* 了解显式和隐式转换。

* See the power of constants and their use in the standard library.
* 查看常量在标准库中的使用。

## Links

[Constants specification](https://golang.org/ref/spec#Constants)    
[Constants](https://blog.golang.org/constants) - Rob Pike    
[Introduction To Numeric Constants In Go](https://www.ardanlabs.com/blog/2014/04/introduction-to-numeric-constants-in-go.html) - William Kennedy

## Code Review

[Declare and initialize constants](example/declare-initalize/declare-init.go)  
[Parallel type system (Kind)](example/parallel-type/parallel-type.go)  
[iota](example/iota/iota.go)  
[Implicit conversion](example/implicit-conversion/implicit-conversion.go)
