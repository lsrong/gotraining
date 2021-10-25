## Variables 变量
Variables are at the heart of the language and provide the ability to read from and write to memory. In Go, access to memory is type safe. This means the compiler takes type seriously and will not allow us to use variables outside the scope of how they are declared.

变量是语言的核心，提供读取和写入内存的能力。在 Go 中，对内存的访问是类型安全的。这意味着编译器会认真对待类型，不允许我们使用超出声明范围的变量

## Notes 重点
* The purpose of all programs and all parts of those programs is to transform data from one form to the other.
* 所有程序及其所有部分的目的是将数据从一种形式转换为另一种形式。
* Code primarily allocates, reads and writes to memory.
* 代码主要分配、读取和写入内存。
* Understanding type is crucial to writing good code and understanding code.
* 理解类型对于编写好的代码和理解代码至关重要。
* If you don't understand the data, you don't understand the problem.
* 如果你不理解数据，你就不会理解问题。
* You understand the problem better by understanding the data.
* 通过了解数据，您可以更好地理解问题。
* When variables are being declared to their zero value, use the keyword `var`.
* 当变量被声明为零值时，使用关键字 `var`。
* When variables are being declared and initialized, use the short variable declaration operator `:=`.
* 在声明和初始化变量时，使用短变量声明运算符`:=`。

## Links 资料链接
[Built-In Types](http://golang.org/ref/spec#Boolean_types)  [内置类型中文文档](https://go-zh.org/ref/spec.old#%E5%B8%83%E5%B0%94%E7%B1%BB%E5%9E%8B)  

[Effective_GO Variables](https://golang.org/doc/effective_go.html#variables)  
Variables can be initialized just like constants but the initializer can be a general expression computed at run time.  
变量可以像常量一样初始化，但初始化器可以是在运行时计算的通用表达式。
```go
var (
	home = os.GetEnv("HOME")
	user = os.GetEnv("USER")
	gopath = os.GetEnv("GOPATH")
)
```

[What's in a name](https://www.youtube.com/watch?v=sFUSP8Au_PE)  (待看)

[A brief history of “type”](http://arcanesentiment.blogspot.com/2015/01/a-brief-history-of-type.html) - Arcane Sentiment  



## Code Review
[Declare and initialize variables](example1/example1.go) ([Go Playground](https://play.golang.org/p/xD_6ghgB7wm))  

[Exercise declare and initialize variables](exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/Ygxt9kW_WAV))
