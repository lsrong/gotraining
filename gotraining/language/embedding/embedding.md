## Embedding

Embedding types provide the final piece of sharing and reusing state and behavior between types. Through the use of inner type promotion, an inner type's fields and methods can be directly accessed by references of the outer type.  
嵌入类型提供了类型之间共享和重用状态和行为的最后一部分。通过使用内部类型提升，外部类型的引用可以直接访问内部类型的字段和方法。
## Notes

* Embedding types allow us to share state or behavior between types.
* 嵌入类型允许我们在类型之间共享状态或行为。
* The inner type never loses its identity.
* 内部类型永远不会失去它的身份。
* This is not inheritance.
* 嵌入并不是继承.
* Through promotion, inner type fields and methods can be accessed through the outer type.
* 通过提升，可以通过外部类型访问内部类型的字段和方法。
* The outer type can override the inner type's behavior.
* 外部类型可以覆盖内部类型的行为

## Links

[Methods, Interfaces and Embedded Types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html) - William Kennedy    
[Embedding is not inheritance](https://rakyll.org/typesystem/) - JBD

## Code Review

[Declaring Fields](example1/example1.go) ([Go Playground](https://play.golang.org/p/mT4iWg10YEp))  
[Embedding types](example2/example2.go) ([Go Playground](https://play.golang.org/p/avo8I21N-qq))  
[Embedded types and interfaces](example3/example3.go) ([Go Playground](https://play.golang.org/p/pdwB9dxD1MR))  
[Outer and inner type interface implementations](example4/example4.go) ([Go Playground](https://play.golang.org/p/soB4QujV4Sj))

## Exercises

### Exercise 1

Copy the code from the template. Add a new type CachingFeed which embeds Feed and overrides the Fetch method.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/_SnIBh3H-0O)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/yHFpf7QYtnc))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
