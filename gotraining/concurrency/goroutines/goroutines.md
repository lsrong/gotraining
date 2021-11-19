""## Goroutines

Goroutines are functions that are created and scheduled to be run independently by the Go scheduler. The Go scheduler is responsible for the management and execution of goroutines.

Goroutines 是由 Go 调度程序创建和调度以独立运行的函数。 Go 调度器负责 goroutine 的管理和执行。

## Notes

* Goroutines are functions that are scheduled to run independently.
* Goroutines 是计划独立运行的函数。
* 
* We must always maintain an account of running goroutines and shutdown cleanly.
* 我们必须始终保持运行 goroutines 并干净地关闭的帐户。
* 
* Concurrency is not parallelism(并行).
    * Concurrency is about dealing with lots of things at once.
    * 并发是关于同时处理很多事情。
    * Parallelism is about doing lots of things at once.
    * 并行是关于同时做很多事情。

_"Parallelism is about physically doing two or more things at the same time. Concurrency is about undefined, out of order, execution." - William Kennedy_    
并行是指在物理上同时做两件事或多件事。 并发是关于未定义的、乱序的、执行的。

_"By default, goroutines shouldn't outlive the function they were created from. this forces you into a extremely good design posture." - Peter Bourgon_  
"默认情况下，goroutines 不应该比创建它们的函数寿命更长。这迫使你进入一个非常好的设计状态。"

## Diagrams

### How the scheduler works.

![Ardan Labs](scheduler.png?v=2)

## Links

[Scheduling In Go - Part I](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html) - William Kennedy    
[Scheduling In Go - Part II](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html) - William Kennedy    
[Scheduler Tracing In Go](https://www.ardanlabs.com/blog/2015/02/scheduler-tracing-in-go.html) - William Kennedy   
[Advanced Go Concurrency Patterns](https://blog.golang.org/advanced-go-concurrency-patterns) - Sameer Ajmani    
[Go Concurrency Patterns: Context](https://blog.golang.org/context) - Sameer Ajmani    
[Concurrency is not parallelism](https://blog.golang.org/concurrency-is-not-parallelism) - Rob Pike    
[Go, for Distributed Systems](https://talks.golang.org/2013/distsys.slide) - Russ Cox    
[Go 1.5 GOMAXPROCS Default](https://docs.google.com/document/d/1At2Ls5_fhJQ59kDK2DFVhFu3g5mATSXqqV5QrxinasI/edit)    
[Concurrency, Goroutines and GOMAXPROCS](https://www.ardanlabs.com/blog/2014/01/concurrency-goroutines-and-gomaxprocs.html) - William Kennedy    
[The Linux Scheduler: a Decade of Wasted Cores](http://www.ece.ubc.ca/~sasha/papers/eurosys16-final29.pdf)    
[Explanation of the Scheduler](https://news.ycombinator.com/item?id=12460807)    
[15 Years of Concurrency](http://joeduffyblog.com/2016/11/30/15-years-of-concurrency/) - Joe Duffy    
[How does the golang scheduler work?](https://www.quora.com/How-does-the-golang-scheduler-work/answer/Ian-Lance-Taylor) - Ian Lance Taylor    
[The Scheduler Saga](https://www.youtube.com/watch?v=YHRO5WQGh0k) - Kavya Joshi

## Code Review

[Goroutines and concurrency](example/goroutine_concurrency/goroutine_concurrency.go) ([Go Playground](https://play.golang.org/p/4n6G3uRDc83))  
[Goroutine time slicing](example2/example2.go) ([Go Playground](https://play.golang.org/p/QtNVo1nb4uQ))  

## Exercises

### Exercise 1

**Part A** Create a program that declares two anonymous functions. One that counts down from 100 to 0 and one that counts up from 0 to 100. Display each number with an unique identifier for each goroutine. Then create goroutines from these functions and don't let main return until the goroutines complete.

**Part B** Run the program in parallel.

[Template](exercises/template1/template1.go) ([Go Playground](https://play.golang.org/p/O0FB2gd6-7d)) |
[Answer](exercises/exercise1/exercise1.go) ([Go Playground](https://play.golang.org/p/uZlHjwf2CXY))
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
""