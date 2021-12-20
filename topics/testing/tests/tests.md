## Testing

Testing is built right into the go tools and the standard library. Testing needs to be a vital part of the development process because it can save you a tremendous amount of time throughout the life cycle of the project. Benchmarking is also a very powerful tool tied to the testing functionality. Aspect of your code can be setup to be benchmarked for performance reviews. Profiling provides a view of the interations between functions and which functions are most heavily used.

测试内置于 go 工具和标准库中。测试需要成为开发过程的重要组成部分，因为它可以在项目的整个生命周期中为您节省大量时间。基准测试也是与测试功能相关的非常强大的工具。可以设置您的代码的某个方面以进行性能审查的基准测试。分析提供了函数之间交互以及哪些函数使用最频繁的视图。

## Notes

* The Go toolset has support for testing and benchmarking.  
  Go 工具集支持测试和基准测试。
  
* The tools are very flexible and give you many options.  
  这些工具非常灵活，可为您提供多种选择。

* Write tests in tandem with development.  
  在开发的同时编写测试。

* Example code serve as both a test and documentation.  
  示例代码既用作测试又用作文档。

* Benchmark throughout the dev, qa and release cycles.  
  整个开发、质量保证和发布周期的基准测试。

* If performance problems are observed, profile your code to see what functions to focus on.  
  如果观察到性能问题，请分析您的代码以查看要关注哪些功能。  
 
* The tools can interfere with each other. For example, precise memory profiling skews CPU profiles, goroutine blocking profiling affects scheduler trace, etc. Rerun tests for each needed profiling mode.  
  这些工具可能会相互干扰。例如，精确的内存分析会扭曲 CPU 配置文件，goroutine 阻塞分析会影响调度程序跟踪等。为每种需要的分析模式重新运行测试。
  
## Quotes

_"A unit test is a test of behavior whose success or failure is wholly determined by the correctness of the test and the correctness of the unit under test." - Kevlin Henney_

## Links

[The deep synergy between testability and good design](https://www.youtube.com/watch?reload=9&feature=share&v=4cVZvoFGJTU&app=desktop) - Michael Feathers  
[testing package](http://golang.org/pkg/testing/)    
[How to write benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go) - Dave Cheney    
[Profiling and creating call graphs for Go programs with "go tool pprof"](http://saml.rilspace.com/profiling-and-creating-call-graphs-for-go-programs-with-go-tool-pprof) - Samuel Lampa    
[pprof package](https://golang.org/pkg/net/http/pprof/)    
[Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs) - Dmitry Vyukov    
https://github.com/dvyukov/go-fuzz  
[Go Dynamic Tools](https://talks.golang.org/2015/dynamic-tools.slide#1) - Dmitry Vyukov    
[Automated Testing with go-fuzz](https://vimeo.com/141698770) - Filippo Valsorda    
[Structuring Tests in Go](https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c#.b2m3nziyb) - Ben Johnson  
[Advanced Testing Concepts for Go 1.7](https://speakerdeck.com/mpvl/advanced-testing-concepts-for-go-1-dot-7) - Marcel van Lohuizen  
[Parallelize your table-driven tests](https://rakyll.org/parallelize-test-tables/) - JBD     
[Advanced Testing with Go - Video](https://www.youtube.com/shared?ci=LARb45o5TpA) - Mitchell Hashimoto  
[Advanced Testing with Go - Deck](https://speakerdeck.com/mitchellh/advanced-testing-with-go) - Mitchell Hashimoto  
[The tragedy of 100% code coverage](http://labs.ig.com/code-coverage-100-percent-tragedy) - Daniel Lebrero's

## Code Review

[Basic Unit Test](example1/example1_test.go) ([Go Playground](https://play.golang.org/p/F7kXmSfr7AE))  
[Table Unit Test](example2/example2_test.go) ([Go Playground](https://play.golang.org/p/1a2u8omEqrX))  
[Mocking Web Server Response](example3/example3_test.go) ([Go Playground](https://play.golang.org/p/SILnu117hak))  
[Testing Internal Endpoints](example4/handlers/handlers_test.go) ([Go Playground](https://play.golang.org/p/CSK7SZEeWf3))  
[Example Test](example4/handlers/handlers_example_test.go) ([Go Playground](https://play.golang.org/p/rE0DRliZH9t))  
[Sub Tests](example5/example5_test.go) ([Go Playground](https://play.golang.org/p/7PrkFU-qVdY))

_Look at the profiling topic to learn more about using test to [profile](../profiling) code._

## Coverage

Making sure your tests cover as much of your code as possible is critical. Go's testing tool allows you to create a profile for the code that is executed during all the tests and see a visual of what is and is not covered.

确保您的测试尽可能多地覆盖您的代码至关重要。 Go 的测试工具允许您为在所有测试期间执行的代码创建配置文件，并查看涵盖和未涵盖的内容。  
```
go test -coverprofile cover.out  
go tool cover -html=cover.out
```

![figure1](testing_coverage.png)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
