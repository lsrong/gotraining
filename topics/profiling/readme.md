## Profiling Code

We can use the go tooling to inspect and profile our programs. Profiling is more of a journey and detective work. It requires some understanding about the application and expectations. The profiling data in and of itself is just raw numbers. We have to give it meaning and understanding.  
我们可以使用 go 工具来检查和分析我们的程序。剖析更像是一次旅行和侦探工作。它需要对应用程序和期望有所了解。分析数据本身只是原始数字。我们必须赋予它意义和理解。

## The Basics of Profiling

_"Those who can make you believe absurdities can make you commit atrocities" - Voltaire_

### How does a profiler work?

A profiler runs your program and configures the operating system to interrupt it at regular intervals. This is done by sending SIGPROF to the program being profiled, which suspends and transfers execution to the profiler. The profiler then grabs the program counter for each executing thread and restarts the program.  
探查器运行您的程序并配置操作系统以定期中断它。这是通过向被分析的程序发送 SIGPROF 来完成的，该程序暂停并将执行转移到分析器。然后分析器获取每个正在执行的线程的程序计数器并重新启动程序。

### Profiling do's and don't's

Before you profile, you must have a stable environment to get repeatable results.

* The machine must be idle—don't profile on shared hardware, don't browse the web while waiting for a long benchmark to run.
* Watch out for power saving and thermal scaling.
* Avoid virtual machines and shared cloud hosting; they are too noisy for consistent measurements.

If you can afford it, buy dedicated performance test hardware. Rack them, disable all the power management and thermal scaling and never update the software on those machines.

For everyone else, have a before and after sample and run them multiple times to get consistent results.

### Types of Profiling

**CPU profiling**  
CPU profiling is the most common type of profile. When CPU profiling is enabled, the runtime will interrupt itself every 10ms and record the stack trace of the currently running goroutines. Once the profile is saved to disk, we can analyse it to determine the hottest code paths. The more times a function appears in the profile, the more time that code path is taking as a percentage of the total runtime.  
CPU 分析是最常见的分析类型。当启用 CPU 分析时，运行时将每 10 毫秒中断一次并记录当前运行的 goroutine 的堆栈跟踪。将配置文件保存到磁盘后，我们可以对其进行分析以确定最热的代码路径。函数在配置文件中出现的次数越多，该代码路径占用的时间占总运行时间的百分比就越多。

**Memory profiling**  
Memory profiling records the stack trace when a heap allocation is made. Memory profiling, like CPU profiling is sample based. By default memory profiling samples 1 in every 1000 allocations. This rate can be changed. Stack allocations are assumed to be free and are not tracked in the memory profile. Because of memory profiling is sample based and because it tracks allocations not use, using memory profiling to determine your application's overall memory usage is difficult.  
当进行堆分配时，内存分析记录堆栈跟踪。内存分析与 CPU 分析一样是基于样本的。默认情况下，内存分析在每 1000 次分配中取样 1。这个比率是可以改变的。假定堆栈分配是空闲的，并且不会在内存配置文件中进行跟踪。因为内存分析是基于样本的，并且因为它跟踪未使用的分配，所以使用内存分析来确定应用程序的整体内存使用情况很困难。

**Block profiling**  
Block profiling is quite unique. A block profile is similar to a CPU profile, but it records the amount of time a goroutine spent waiting for a shared resource. This can be useful for determining concurrency bottlenecks in your application. Block profiling can show you when a large number of goroutines could make progress, but were blocked.  
阻塞分析非常独特。块配置文件类似于 CPU 配置文件，但它记录了 goroutine 等待共享资源所花费的时间。这对于确定应用程序中的并发瓶颈很有用。块分析可以显示大量 goroutine 何时可以取得进展但被阻止。

Blocking includes:
阻塞包括：

* Sending or receiving on an unbuffered channel.  
  在无缓冲通道上发送或接收。
* Sending to a full channel, receiving from an empty one.  
  发送到一个完整的channel，从一个空的频道接收。
* Trying to Lock a sync.Mutex that is locked by another goroutine.  
  试图锁定被另一个 goroutine 锁定的 sync.Mutex。
* Block profiling is a very specialised tool, it should not be used until you believe you have eliminated all your CPU and memory usage bottlenecks.
  阻塞分析是一种非常专业的工具，在您确信已消除所有 CPU 和内存使用瓶颈之前，不应使用它。

**One profile at time**  
**一次启用一个分析**  
Profiling is not free. Profiling has a moderate, but measurable impact on program performance—especially if you increase the memory profile sample rate. Most tools will not stop you from enabling multiple profiles at once. If you enable multiple profiles at the same time, they will observe their own interactions and skew your results.  
分析不是免费的。分析对程序性能有中等但可衡量的影响——尤其是当您提高内存分析采样率时。大多数工具不会阻止您一次启用多个配置文件。如果您同时启用多个配置文件，它们将观察自己的交互并扭曲您的结果。

**Do not enable more than one kind of profile at a time.**  
**不要一次启用多种配置文件。**

### Hints to interpret what you see in the profile

If you see lots of time spent in `runtime.mallocgc` function, the program potentially makes excessive amount of small memory allocations. The profile will tell you where the allocations are coming from. See the memory profiler section for suggestions on how to optimize this case.  
如果您看到在 `runtime.mallocgc` 函数中花费了大量时间，则该程序可能会进行过多的小内存分配。配置文件将告诉您分配的来源。有关如何优化这种情况的建议，请参阅内存分析器部分。

If lots of time is spent in channel operations, `sync.Mutex` code and other synchronization primitives or System component, the program probably suffers from contention. Consider to restructure program to eliminate frequently accessed shared resources. Common techniques for this include sharding/partitioning, local buffering/batching and copy-on-write technique.  
如果在通道操作、`sync.Mutex` 代码和其他同步原语或系统组件上花费了大量时间，则程序可能会出现争用。考虑重构程序以消除频繁访问的共享资源。常用的技术包括分片分区、本地缓冲批处理和写时复制技术。

If lots of time is spent in `syscall.Read/Write`, the program potentially makes excessive amount of small reads and writes. Bufio wrappers around os.File or net.Conn can help in this case.  
如果在 `syscall.ReadWrite` 上花费大量时间，程序可能会进行过多的小读和写操作。在这种情况下，围绕 os.File 或 net.Conn 的 Bufio 包装器可以提供帮助。

If lots of time is spent in GC component, the program either allocates too many transient objects or heap size is very small so garbage collections happen too frequently.  
如果在 GC 组件上花费大量时间，则程序要么分配过多的临时对象，要么堆大小非常小，因此垃圾收集发生得太频繁。

* Large objects affect memory consumption and GC time, while large number of tiny allocations affects execution speed.    
  大对象会影响内存消耗和 GC 时间，而大量的微小分配会影响执行速度。

* Combine values into larger values. This will reduce number of memory allocations (faster) and also reduce pressure on garbage collector (faster garbage collections).  
  将值组合成更大的值。这将减少内存分配的数量（更快），也减少垃圾收集器的压力（更快的垃圾收集）

* Values that do not contain any pointers are not scanned by garbage collector. Removing pointers from actively used value can positively impact garbage collection time.  
  垃圾收集器不会扫描不包含任何指针的值。从积极使用的值中删除指针可以对垃圾收集时间产生积极影响。

## Rules of Performance

1) Never guess about performance.   永远不要猜测性能。
2) Measurements must be relevant.   测量必须是相关的。
3) Profile before you decide something is performance critical.     在您决定某些事情对性能至关重要之前进行概要分析。
4) Test to know you are correct.    测试以确认您是正确的

## Installing Tools

**hey**  
hey is a modern HTTP benchmarking tool capable of generating the load you need to run tests. It's built using the Go language and leverages goroutines for behind the scenes async IO and concurrency.
hey 是一种现代 HTTP 基准测试工具，能够生成运行测试所需的负载。它是使用 Go 语言构建的，并利用 goroutines 在幕后实现异步 IO 和并发。

    go get -u github.com/rakyll/hey

## Dave Cheney's Profiling Presentation:

Much of what I have learned comes from Dave and working on solving problems. This slide deck is a great place to start. Much of this material can be found in the material below.

[Seven ways to profile a Go program](http://go-talks.appspot.com/github.com/davecheney/presentations/seven.slide#1)

## Profiling, Debugging and Optimization Reading

Here is more reading and videos to also help get you started.

[The past and future of Microprocessor performance](https://github.com/davecheney/gophercon2018-performance-tuning-workshop/blob/master/1-welcome/introduction.md) - Dave Cheney

[Language Mechanics On Escape Analysis](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-escape-analysis.html) - William Kennedy  
[Profiling Go Programs](http://golang.org/blog/profiling-go-programs) - Go Team  
[Profiling & Optimizing in Go](https://www.youtube.com/watch?v=xxDZuPEgbBU) - Brad Fitzpatrick  
[Go Dynamic Tools](https://www.youtube.com/watch?v=a9xrxRsIbSU) - Dmitry Vyukov  
[How NOT to Measure Latency](https://www.youtube.com/watch?v=lJ8ydIuPFeU&feature=youtu.be) - Gil Tene  
[Go Performance Tales](http://jmoiron.net/blog/go-performance-tales) - Jason Moiron  
[Debugging performance issues in Go programs](https://software.intel.com/en-us/blogs/2014/05/10/debugging-performance-issues-in-go-programs) - Dmitry Vyukov  
[Reduce allocation in Go code](https://methane.github.io/2015/02/reduce-allocation-in-go-code) - Python Bytes  
[Write High Performance Go](http://go-talks.appspot.com/github.com/davecheney/presentations/writing-high-performance-go.slide) - Dave Cheney  
[Static analysis features of godoc](https://golang.org/lib/godoc/analysis/help.html) - Go Team   
[Seven ways to profile a Go program](https://www.bigmarker.com/remote-meetup-go/Seven-ways-to-profile-a-Go-program) - Dave Cheney   
[runtime: goroutine execution stalled during GC](https://github.com/golang/go/issues/16293) - Caleb Spare  
[Go's execution tracer](http://www.thedotpost.com/2016/10/rhys-hiltner-go-execution-tracer) - Rhys Hiltner  
[Using Instruments to profile Go programs](https://rakyll.org/instruments) - JBD    
[Fighting latency: the CPU profiler is not your ally](https://www.youtube.com/watch?v=nsM_m4hZ-bA&t=973s) - Filippo Valsorda  
[go tool trace](https://making.pusher.com/go-tool-trace/) - Will Sewell

## Go and OS Tooling

### time

The **time** command provide information that can help you get a sense how your program is performing.  
time 命令提供的信息可以帮助您了解程序的执行情况。

Use the **time** command to see data about building the program.  
$ cd $GOPATH/src/github.com/ardanlabs/gotraining/topics/go/profiling/project  
$ /usr/bin/time -lp go build		-- Mac OS X  
$ /usr/bin/time -v go build		-- Linux  

### perf

If you're a linux user, then perf(1) is a great tool for profiling applications. Now we have frame pointers, perf can profile Go applications.

	$ go build -toolexec="perf stat" cmd/compile/internal/gc
	# cmd/compile/internal/gc

 	Performance counter stats for '/home/dfc/go/pkg/tool/linux_amd64/compile -o $WORK/cmd/compile/internal/gc.a -trimpath $WORK -p cmd/compile/internal/gc -complete -buildid 87cd803267511b4d9e753d68b5b66a70e2f878c4 -D _/home/dfc/go/src/cmd/compile/internal/gc -I $WORK -pack ./alg.go ./align.go ./bexport.go ./bimport.go ./builtin.go ./bv.go ./cgen.go ./closure.go ./const.go ./cplx.go ./dcl.go ./esc.go ./export.go ./fmt.go ./gen.go ./go.go ./gsubr.go ./init.go ./inl.go ./lex.go ./magic.go ./main.go ./mpfloat.go ./mpint.go ./obj.go ./opnames.go ./order.go ./parser.go ./pgen.go ./plive.go ./popt.go ./racewalk.go ./range.go ./reflect.go ./reg.go ./select.go ./sinit.go ./sparselocatephifunctions.go ./ssa.go ./subr.go ./swt.go ./syntax.go ./type.go ./typecheck.go ./universe.go ./unsafe.go ./util.go ./walk.go':

       7026.140760 task-clock (msec)         #    1.283 CPUs utilized          
             1,665 context-switches          #    0.237 K/sec                  
                39 cpu-migrations            #    0.006 K/sec                  
            77,362 page-faults               #    0.011 M/sec                  
    21,769,537,949 cycles                    #    3.098 GHz                     [83.41%]
    11,671,235,864 stalled-cycles-frontend   #   53.61% frontend cycles idle    [83.31%]
     6,839,727,058 stalled-cycles-backend    #   31.42% backend  cycles idle    [66.65%]
    27,157,950,447 instructions              #    1.25  insns per cycle        
                                             #    0.43  stalled cycles per insn [83.25%]
     5,351,057,260 branches                  #  761.593 M/sec                   [83.49%]
       118,150,150 branch-misses             #    2.21% of all branches         [83.15%]

       5.476816754 seconds time elapsed

## Basic Go Profiling

Learn the basics of reading Stack Traces.  
[Stack Traces and Core Dumps](stack_trace/README.md)

Learn the basics of using GODEBUG.  
[GODEBUG](godebug/README.md)

Learn the basics of using memory and cpu profiling.  
[Memory and CPU Profiling](memcpu/README.md)

Learn the basics of using http/pprof.  
[pprof Profiling](pprof/README.md)

Learn the basics of blocking profiling.  
[Blocking Profiling](blocking/README.md)

Learn the basics of mutex profiling.  
[Mutex Profiling](mutex/README.md)

Learn the basics of tracing.    
[Tracing](trace/README.md)

Learn the basics of profiling and tracing a larger application.  
[Real World Example](project/README.md)

## Godoc Analysis

The `godoc` tool can help you perform static analysis on your code.

	// Perform a pointer analysis and then run the godoc website.
	$ godoc -analysis pointer -http=:8080

[Static analysis features of godoc](https://golang.org/lib/godoc/analysis/help.html) - Go Team

## HTTP Tracing

HTTP tracing facilitate the gathering of fine-grained information throughout the lifecycle of an HTTP client request.

[HTTP Tracing Package](http_trace/README.md)

___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
