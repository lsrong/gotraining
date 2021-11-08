## Maps

Maps provide a data structure that allow for the storage and management of key/value pair data.
Maps 提供了一种数据结构，允许存储和管理键值对数据。

## Notes

* Maps provide a way to store and retrieve key/value pairs.
* Maps 提供了一种存储和检索键值对的方法。

* Reading an absent key returns the zero value for the map's value type.
* 读取不存在的键会返回映射值类型的零值。

* Iterating over a map is always random.
* 迭代map总是随机的。

* The map key must be a value that is comparable.
* 映射键必须是可比较的值。

* Elements in a map are not addressable.
* 映射中的元素不可寻址。

* Maps are a reference type.
* map是一种引用类型。

## Links

[Go maps in action](https://blog.golang.org/go-maps-in-action) - Andrew Gerrand    
[Macro View of Map Internals In Go](https://www.ardanlabs.com/blog/2013/12/macro-view-of-map-internals-in-go.html) - William Kennedy    
[Inside the Map Implementation](https://www.youtube.com/watch?v=Tl7mi9QmLns) - Keith Randall    
[How the Go runtime implements maps efficiently (without generics)](https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics) - Dave Cheney

## Code Review

[Declare, write, read, and delete](example/declaration/declaration.go)  
[Absent keys](example/absent/absent.go)  
[Map key restrictions](example/restriction/restriction.go)  
[Map literals and range](example/iterate/iterate.go)  
[Sorting maps by key](example/sorting/sorting.go)  
[Taking an element's address](example/not-addressable/not-addressable.go)  
[Maps are Reference Types](example/reference/reference.go)

## Exercises

### Exercise 1

Declare and make a map of integer values with a string as the key. Populate the map with five values and iterate over the map to display the key/value pairs.
声明并制作一个以字符串为键的整数值映射。用五个值填充地图并迭代地图以显示键值对

[Exercise](exercise/exercise.go)