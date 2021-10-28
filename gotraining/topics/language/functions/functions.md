## Functions

Functions are at the core of the language. They provide a mechanism to group and organize our code to separate and distinct pieces of functionality. They can be used to provide an API to the packages we write and are a core component to concurrency.  
函数是语言的核心。它们提供了一种机制来对我们的代码进行分组和组织，以分离和区分不同的功能。它们可用于为我们编写的包提供 API，并且是并发的核心组件。
## Notes

* Functions can return multiple values and most return an error value.
* 函数可以返回多个值并且大多数返回错误值。

* The error value should always be checked as part of the programming logic.
* 错误值应始终作为编程逻辑的一部分进行检查。

* The blank identifier can be used to ignore return values.
* 空白标识符可用于忽略返回值。

## Links

[Functions](https://golang.org/doc/effective_go.html#functions)  
[Functions-and-naked-returns-in-go](https://www.ardanlabs.com/blog/2013/10/functions-and-naked-returns-in-go.html)  
[Understanding-defer-panic-and-recover](https://www.ardanlabs.com/blog/2013/06/understanding-defer-panic-and-recover.html)

## Code Review

[Return multiple values](example1/example1.go) ([Go Playground](https://play.golang.org/p/-7A-lGLv2TK))  
[Blank identifier](example2/example2.go) ([Go Playground](https://play.golang.org/p/ID54tVxM5B0))  
[Redeclarations](example3/example3.go) ([Go Playground](https://play.golang.org/p/EDRhDh2r1Mj))  
[Anonymous Functions/Closures](example4/example4.go) ([Go Playground](https://play.golang.org/p/h8Yi_2Sxsmu))

## Advanced Code Review

[Recover panics](advanced/example1/example1.go) ([Go Playground](https://play.golang.org/p/Wcd3CLbQZBH))

## Exercises

### Exercise 1

**Part A** Declare a struct type to maintain information about a user. Declare a function that creates value of and returns pointers of this type and an error value. Call this function from main and display the value.

**Part B** Make a second call to your function but this time ignore the value and just test the error value.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/5vEQxEzq3i_D)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/P8wC324WWuh))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
