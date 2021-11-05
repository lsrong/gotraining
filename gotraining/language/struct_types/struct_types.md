## Struct Types 结构体类型

Struct types are a way of creating complex types that group fields of data together. They are a great way of organizing and sharing the different aspects of the data your program consumes.
结构类型是一种创建将数据字段组合在一起的复杂类型的方法。它们是组织和共享程序使用的数据的不同方面的好方法。  

A computer architecture’s potential performance is determined predominantly by its word length (the number of bits that can be processed per access) and, more importantly, memory size, or the number of words that it can access.
计算机体系结构的潜在性能主要取决于其字长（每次访问可以处理的位数），更重要的是，内存大小或可以访问的字数。

## Notes
* We can use the struct literal form to initialize a value from a struct type.
* 我们可以使用结构体字面量形式来初始化结构体类型的值。
* The dot (`.`) operator allows us to access individual field values.
* 点 (`.`) 运算符允许我们访问单个字段值。
* We can create anonymous structs.
* 我们可以创建匿名结构。

## Links

[Understanding Type in Go](https://www.ardanlabs.com/blog/2013/07/understanding-type-in-go.html) - William Kennedy    
[Object Oriented Programming in Go](https://www.ardanlabs.com/blog/2013/07/object-oriented-programming-in-go.html) - William Kennedy    
[Padding is hard](https://dave.cheney.net/2015/10/09/padding-is-hard) - Dave Cheney    
[Structure Member Alignment, Padding and Data Packing](https://www.geeksforgeeks.org/structure-member-alignment-padding-and-data-packing/)    
[The Lost Art of Structure Packing](http://www.catb.org/esr/structure-packing) - Eric S. Raymond

## Code Review

[Declare, create and initialize struct types](example2/example2.go) ([Go Playground](https://play.golang.org/p/djzGT1JtSwy))  
[Anonymous struct types](example3/example3.go) ([Go Playground](https://play.golang.org/p/09cxjnmfcdC))  
[Named vs Unnamed types](example4/example4.go) ([Go Playground](https://play.golang.org/p/ky91roJDjir))

## Advanced Code Review

[Struct type alignments](example1/example1.go) ([Go Playground](https://play.golang.org/p/rAvtS7cgD0z))

## Exercises

### Exercise 1

**Part A:** Declare a struct type to maintain information about a user (name, email and age). Create a value of this type, initialize with values and display each field.

**Part B:** Declare and initialize an anonymous struct type with the same three fields. Display the value.

[Answer](exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/eT_gLZKeHk-))
