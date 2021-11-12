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

[Declaring Fields](example/not-embedding/not-embedding.go)  
[Embedding types](example/embedding-type/embedding-type.go)  
[Embedded types and interfaces](example/embedding-interface/embedding-interface.go)  
[Outer and inner type interface implementations](example/inner-outer-implemention/inner_outer_implemention.go)

## Exercises

### Exercise 1

Copy the code from the template. Add a new type CachingFeed which embeds Feed and overrides the Fetch method.

[Template](template/template.go) |
[Answer](exercise/exercise.go)

