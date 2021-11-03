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

[Return multiple values](example/multiple-values/multiple-values.go)  
[Blank identifier](example/blank-identifier/blank-identifier.go)  
[Redeclarations](example/redeclarations/redeclarations.go)  
[Anonymous Functions/Closures](example/anonymous-functions-clourses/anonymous-clourses.go)
