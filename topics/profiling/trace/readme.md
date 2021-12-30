## Tracing

The tracing can help identify not only what is happening but also what is not happening when your program is running. We will use a simple program to learn how to navigate and read some of the tracing information you can find in the trace tool.  

跟踪不仅可以帮助识别正在发生的事情，还可以帮助识别程序运行时没有发生的事情。我们将使用一个简单的程序来学习如何导航和阅读您可以在跟踪工具中找到的一些跟踪信息。

## Basic Skills

Review this post to gain basic skills.

[go tool trace](https://making.pusher.com/go-tool-trace/) - Will Sewell  
[Debugging Latency in Go 1.11](https://medium.com/observability/debugging-latency-in-go-1-11-9f97a7910d68) - JBD

## Trace Command

You have two options with this code. First uncomment the CPU profile lines to generate a CPU profile.  
使用此代码您有两个选择。首先取消对 CPU 配置文件行的注释以生成 CPU 配置文件。

    pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()

	// trace.Start(os.Stdout)
	// defer trace.Stop()

This will let you run a profile first. Leverage the lessons learned in the other sections.  
这将让您首先运行配置文件。利用在其他部分中学到的经验教训。

    $ ./trace > p.out
    $ go tool pprof p.out

Then run a trace by uncommenting the other lines of code.  
然后通过取消注释其他代码行来运行跟踪。

    // pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()

	trace.Start(os.Stdout)
	defer trace.Stop()

Once you run the program.  
一旦你运行程序。  

    $ ./trace > t.out
    $ go tool trace t.out

Then explore the trace tooling by building the program with these different find functions.  
然后通过使用这些不同的查找功能构建程序来探索跟踪工具。

    n := find(topic, docs)
	// n := findConcurrent(topic, docs)
	// n := findConcurrentSem(topic, docs)
	// n := findProcessors(topic, docs)
	// n := findActor(topic, docs)

Using this function allows you to see how to add custom tasks and regions. This requires Go version 1.11.  
使用此功能可以查看如何添加自定义任务和区域。这需要 Go 1.11 版。

	// n := findProcessorsTasks(topic, docs)

_Note that goroutines in "syscall" state consume an OS thread, other goroutines do not (except for goroutines that called runtime.LockOSThread, which is, unfortunately, not visible in the profile)._
_请注意，处于“系统调用”状态的 goroutine 会消耗 OS 线程，而其他 goroutine 不会（调用 runtime.LockOSThread 的 goroutine 除外，不幸的是，它在配置文件中不可见）。_

_Note that goroutines in "IO wait" state do NOT consume an OS thread. They are parked on the non-blocking network poller._  
_注意处于“IO 等待”状态的 goroutine 不消耗 OS 线程。它们停在非阻塞网络轮询器上。_

## Code Review

[Profiling Test](trace.go)