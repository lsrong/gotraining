## Methods

Methods are functions that give data the ability to exhibit behavior.
方法是赋予数据表现行为能力的函数。

## Notes

* Methods are functions that declare a receiver variable.
* 方法是声明接收者变量的函数

* Receivers bind a method to a type and can use value or pointer semantics.
* 接收者将方法绑定到类型，并且可以使用值或指针语义。

* Value semantics mean a copy of the value is passed across program boundaries.
* 值语义意味着值的副本跨程序边界传递. value.Method() == (T.Method(value) || T.Method(*pointer))

* Pointer semantics mean a copy of the values address is passed across program boundaries.
* 指针语义意味着值地址的副本跨程序边界传递. pointer.Method() == (*T.Method(&value) || *T.Method(pointer))

* Stick to a single semantic for a given type and be consistent.
* 坚持给定类型的单一语义并保持一致。

## Quotes

_"Methods are valid when it is practical or reasonable for a piece of data to expose a capability." - William Kennedy_

## Links

[Methods](https://golang.org/doc/effective_go.html#methods)    
[Methods, Interfaces and Embedded Types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy    
[Escape-Analysis Flaws](https://www.ardanlabs.com/blog/2018/01/escape-analysis-flaws.html) - William Kennedy

## Code Review

[Declare and receiver behavior](example/declaration/declaration.go)  
[Named typed methods](example/type-method/type-method.go)  
[Function/Method variables](example/func-method-variables/func-method-variables.go)  
[Function Types](example/function-type/function-type.go)

## Exercises

Declare a struct that represents a baseball player. Include name, atBats and hits. Declare a method that calculates a players batting average. The formula is Hits / AtBats. Declare a slice of this type and initialize the slice with several players. Iterate over the slice displaying the players name and batting average.
声明一个代表棒球运动员的结构体。包括名称、atBats 和点击数。
声明一种计算球员击球平均值的方法。 公式是Hits / AtBats。
声明一个这种类型的切片并用多个玩家初始化切片。遍历显示球员姓名和击球率的切片。

[Exercise](exercise/exercise.go)
