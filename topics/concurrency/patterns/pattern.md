## Concurrency Patterns
There are lots of different patterns we can create with goroutines and channels. Two interesting patterns are resource pooling and concurrent searching.

我们可以使用 goroutine 和通道创建许多不同的模式。两个有趣的模式是资源池和并发搜索。

## Notes

* The work code provides a pattern for giving work to a set number of goroutines without losing the guarantee.
* 工作代码提供了一种模式，可以在不失去保证的情况下将工作分配给一定数量的 goroutine。
*
* The resource pooling code provides a pattern for managing resources that goroutines may need to acquire and release.
*  资源池代码提供了一种管理 goroutine 可能需要获取和释放的资源的模式。
*
* The search code provides a pattern for using multiple goroutines to perform concurrent work.
*  搜索代码提供了一种使用多个 goroutine 来执行并发工作的模式。

## Links

[Concurrency patterns](https://github.com/gobridge/concurrency-patterns)    
[Go Concurrency Patterns: Pipelines and cancellation](https://blog.golang.org/pipelines) - Sameer Ajmani    
[Go Concurrency Patterns](https://talks.golang.org/2012/concurrency.slide#1) - Rob Pike    
[Go Concurrency Patterns: Context](https://blog.golang.org/context) - Sameer Ajmani    
[Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns) - Sameer Ajmani    
[Go: code that grows with grace](https://talks.golang.org/2012/chat.slide) - Andrew Gerrand

Functional Options : type DialOption func(*dialOptions)  
https://github.com/grpc/grpc-go/blob/master/clientconn.go

## Code Review

[Chat](chat)  
[Logger](logger)  
[Task](task)  
[Pooling](pool)  
[Kit](https://github.com/ardanlabs/kit)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
