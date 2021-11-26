## Exporting

Packages contain the basic unit of compiled code. They define a scope for the identifiers that are declared within them. Exporting is not the same as public and private semantics in other languages. But exporting is how we provide encapsulation in Go.

扩展包含编译代码的基本单元。它们为在其中声明的标识符定义了一个范围。导出与其他语言中的公共和私有语义不同。但是导出是我们在 Go 中提供封装的方式。

## Notes

* Code in go is compiled into packages and then linked together.
* go中的代码被编译成包，然后链接在一起。
* 
* Identifiers are exported (or remain unexported) based on letter-case.
* 标识符基于字母大小写导出（或保持未导出）。
* 
* We import packages to access exported identifiers.
* 我们导入包以访问导出的标识符。
* 
* Any package can use a value of an unexported type, but this is annoying to use.
* 任何包都可以使用未导出类型的值，但这使用起来很烦人。

## Links

[Exported/Unexported Identifiers In Go](https://www.ardanlabs.com/blog/2014/03/exportedunexported-identifiers-in-go.html) - William Kennedy

## Code Review

[Declare and access exported identifiers - Pkg](example/example1/counters/counters.go) ([Go Playground](https://play.golang.org/p/8Xzq-m9ez-I))  
[Declare and access exported identifiers - Main](example/example1/example1.go) ([Go Playground](https://play.golang.org/p/KrpX0CyIyYO))

[Declare unexported identifiers and restrictions - Pkg](example/example2/counters/counters.go) ([Go Playground](https://play.golang.org/p/9u1IQexx5gk))  
[Declare unexported identifiers and restrictions - Main](example/example2/example2.go) ([Go Playground](https://play.golang.org/p/A5FpmRpuOWJ))

[Access values of unexported identifiers - Pkg](example/example3/counters/counters.go) ([Go Playground](https://play.golang.org/p/NroO30yoNvh))  
[Access values of unexported identifiers - Main](example/example3/example3.go) ([Go Playground](https://play.golang.org/p/e5fg0uOEkkn))

[Unexported struct type fields - Pkg](example/example4/users/users.go) ([Go Playground](https://play.golang.org/p/KQ6x5z7E1pN))  
[Unexported struct type fields - Main](example/example4/example4.go) ([Go Playground](https://play.golang.org/p/6MznWaiGwr-))

[Unexported embedded types - Pkg](example/example5/users/users.go) ([Go Playground](https://play.golang.org/p/br-2rVc1VF1))  
[Unexported embedded types - Main](example/example5/example5.go) ([Go Playground](https://play.golang.org/p/p9pQo5gCB42))

## Exercises

### Exercise 1
**Part A** Create a package named toy with a single exported struct type named Toy. Add the exported fields Name and Weight. Then add two unexported fields named onHand and sold. Declare a factory function called New to create values of type toy and accept parameters for the exported fields. Then declare methods that return and update values for the unexported fields.

**Part B** Create a program that imports the toy package. Use the New function to create a value of type toy. Then use the methods to set the counts and display the field values of that toy value.

[Template](exercises/template1) |
[Answer](exercises/exercise1)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
