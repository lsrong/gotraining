## Context - Standard Library

The package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes.

## Notes

* Incoming requests to a server should create a Context.
* 对服务器的传入请求应创建一个上下文。
* 
* Outgoing calls to servers should accept a Context.
* 对服务器的传出调用应该接受一个上下文。
* 
* The chain of function calls between them must propagate the Context.p
* 它们之间的函数调用链必须传播 Context.p
* 
* Replace a Context using WithCancel, WithDeadline, WithTimeout, or WithValue.
* 使用 WithCancel、WithDeadline、WithTimeout 或 WithValue 替换上下文。
* 
* When a Context is canceled, all Contexts derived from it are also canceled.
* 当一个上下文被取消时，所有从它派生的上下文也被取消。
* 
* Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function that needs it.
* 不要将上下文存储在结构类型中；相反，将 Context 显式传递给需要它的每个函数。
* 
* Do not pass a nil Context, even if a function permits it. Pass context.TODO if you are unsure about which Context to use.
* 即使函数允许，也不要传递 nil 上下文。如果不确定要使用哪个 Context，请传递 context.TODO。
* 
* Use context Values only for request-scoped data that transits processes and APIs, not for passing optional parameters to functions.
* 仅将上下文值用于传输流程和 API 的请求范围数据，而不用于将可选参数传递给函数。
* 
* The same Context may be passed to functions running in different goroutines; Contexts are safe for simultaneous use by multiple goroutines.
* 相同的 Context 可以传递给在不同的 goroutine 中运行的函数；上下文对于多个 goroutine 同时使用是安全的。
* 
## Links

[Package context](https://golang.org/pkg/context)  
[Go Concurrency Patterns: Context](https://blog.golang.org/context) - Sameer Ajmani    
[Cancellation, Context, and Plumbing](https://vimeo.com/115309491) - Sameer Ajmani    
[Using contexts to avoid leaking goroutines](https://rakyll.org/leakingctx/) - JBD

## Code Review

**_"Context values are for request-scoped data that passes through programs in a distributed system. Litmus test: Could it be an HTTP header?" - Sameer Ajmani_**

[Leak Goroutine](example/leaking/leaking.go)   
[Store / Retrieve context values](example/example1/example1.go)   
[WithCancel](example/example2/example2.go)   
[WithDeadline](example/example3/example3.go)   
[WithTimeout](example/example4/example4.go)   
[Request/Response](example/example5/example5.go)  
[Cancellation](example/example6/example6.go) 
## Exercises

### Exercise 1

Use the template and follow the directions. You will be writing a web handler that performs a mock database call but will timeout based on a context if the call takes too long. You will also save state into the context.

[Template](template/template.go) ([Go Playground](https://play.golang.org/p/jIkgYBhqMNy)) |
[Answer](exercise/exercise.go)
