## Profiling a Larger Web Service

We have a web application that extends a web service. Let's profile this application and attempt to understand how it is working.  
有一个扩展 Web 服务的 Web 应用程序。让我们分析这个应用程序并尝试了解它是如何工作的。


### Building and Running the Project

We have a website that we will use to learn and explore more about profiling. This project is a search engine for RSS feeds.  
我们有一个网站，我们将用它来学习和探索更多关于分析的信息。这个项目是一个RSS提要的搜索引擎。  

目前对象进行了重写，新闻资讯的搜索引擎，目前只对接了网易新闻。

Run the website and validate it is working.  
运行网站并验证它是否正常工作。

	$ go build
	$ ./project

	http://localhost:4000

### Adding Load

To add load to the service while running profiling we can run these command.  
要在运行分析时向服务添加负载，我们可以运行这些命令 `hey HTTP 基准测试工具，安装命令： go get -u github.com/rakyll/hey`

	// Send 10k request using 100 connections.
	$ hey -m POST -c 100 -n 10000 "http://localhost:4000/search?keyword=人民&cm=on"

### GODEBUG

#### GC Trace

Run the website redirecting stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.  
运行将标准输出（日志）重定向到空设备的网站。这将允许我们只看到来自运行时的跟踪信息。

	$ GODEBUG=gctrace=1 ./project > /dev/null

```
gc # @#s #%: #+#+# ms clock, #+#/#/#+# ms cpu, #->#-># MB, # MB goal, # P
where the fields are as follows:
gc #        the GC number, incremented at each GC, GC编号
@#s         time in seconds since program start， 程序启动后的时间（以秒为单位）
#%          percentage of time spent in GC since program start， 自程序启动以来在 GC 上花费的时间百分比
#+...+#     wall-clock/CPU times for the phases of the GC， GC 阶段的挂钟 CPU 时间
#->#-># MB  heap size at GC start, at GC end, and live heap， GC 开始时、GC 结束时和活动堆时的堆大小
# MB goal   goal heap size， 目标堆大小
# P         number of processors used，使用的处理器数量

eg:
gc 2814 @39.333s 1%: 0.26+1.9+0.91 ms clock, 2.1+0.11/0.53/0+7.3 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
2814 表示第2814次执行
@39.333s 表示程序执行的总时间
1% 垃圾回收时间占用的百分比
0.26+1.9+0.91 ms clock 垃圾回收的时间，分别为STW（stop-the-world）清扫的时间, 并发标记和扫描的时间，STW标记的时间
2.1+0.11/0.53/0+7.3 ms cpu 垃圾回收占用cpu时间
4->4->1 MB 堆的大小，gc后堆的大小，存活堆的大小
5 MB goal 整体堆的大小
8 P 使用的处理器数量
```


#### GOGC

GOGC will change the way the heap grows. Changing this value could help reduce the number of GC's that occur.  
GOGC 将改变堆的增长方式。更改此值有助于减少发生的 GC 次数。

Run the website again adding load. Look at the pacing of the GC with these different GOGC values.  
再次运行网站增加负载。查看具有这些不同 GOGC 值的 GC 的步调。

	$ GODEBUG=gctrace=1 ./project > /dev/null  
	$ GODEBUG=gctrace=1 GOGC=200 ./project > /dev/null  
	$ GODEBUG=gctrace=1 GOGC=500 ./project > /dev/null

#### Scheduler Trace

Run the website redirecting stdout (logs) to the null device. This will allow us to just see the trace information from the runtime.   
运行将标准输出（日志）重定向到空设备的网站。这将允许我们只看到来自运行时的跟踪信息。  

	$ GODEBUG=schedtrace=1000 ./project > /dev/null

追踪结果显示：
```
SCHED 126339ms: gomaxprocs=8 idleprocs=0 threads=17 spinningthreads=1 idlethreads=4 runqueue=0 [2 0 0 0 0 0 0 0]
上面输出的说明：
SCHED 126339ms: 启动到现在输出行的运行时间，126339ms
gomaxprocs=8: 当前运行8个cpu核心
idleprocs=0: 空闲核心数为0
threads=17：前正在运行的OS线程数，17
spinningthreads=1: 自旋状态的 OS 线程数量。自旋锁（Spin Lock）是一种简单、高效、线程安全的同步原语（synchronization primitive），其在等待时会反复检查锁的状态，直到解锁。
idlethreads=4：空闲的线程数量。
runqueue=0 [2 0 0 0 0 0 0 0]：全局队列中的 Goroutine 数量，而后面的 [2 0 0 0 0 0 0 0] 则分别代表第 n 个 P 的本地队列正在运行的 Goroutine 数量。
```

### PPROF

We already added the following import so we can include the profiling route to our web service.  
我们已经添加了以下导入，因此我们可以将分析路由包含到我们的 Web 服务中。

	import _ "net/http/pprof"

#### Raw http/pprof

Look at the basic profiling stats from the new endpoint:  
查看来自新端点的基本分析统计信息：

	http://localhost:4000/debug/pprof

Capture heap profile:  
捕获堆信息文件：

	http://localhost:4000/debug/pprof/heap

Capture cpu profile:  
捕获 CPU 信息文件：

	http://localhost:4000/debug/pprof/profile

#### Interactive Profiling

Run the Go pprof tool in another window or tab to review alloc space heap information.  
在另一个窗口或选项卡中运行 Go pprof 工具以查看分配空间堆信息。

	$ go tool pprof http://localhost:4000/debug/pprof/allocs

Documentation of memory profile options.  
内存配置文件选项的文档。

    // Useful to see current status of heap.
	-inuse_space  : Allocations live at the time of profile  	** default
	-inuse_objects: Number of bytes allocated at the time of profile

	// Useful to see pressure on heap over time.
	-alloc_space  : All allocations happened since program start
	-alloc_objects: Number of object allocated at the time of profile

If you want to reduce memory consumption, look at the `-inuse_space` profile collected during normal program operation.  
如果要减少内存消耗，请查看在正常程序运行期间收集的`-inuse_space` 配置文件。

If you want to improve execution speed, look at the `-alloc_objects` profile collected after significant running time or at program end.  
如果您想提高执行速度，请查看在大量运行时间或程序结束后收集的 `-alloc_objects` 配置文件。

Run the Go pprof tool in another window or tab to review cpu information.  
在另一个窗口或选项卡中运行 Go pprof 工具以查看 cpu 信息。

	$ go tool pprof http://localhost:4000/debug/pprof/profile

_Note that goroutines in "syscall" state consume an OS thread, other goroutines do not (except for goroutines that called runtime.LockOSThread, which is, unfortunately, not visible in the profile)._  
_请注意，处于“系统调用”状态的 goroutine 会消耗 OS 线程，而其他 goroutine 不会（调用 runtime.LockOSThread 的 goroutine 除外，不幸的是，它在配置文件中不可见）。_   

_Note that goroutines in "IO wait" state do NOT consume an OS thread. They are parked on the non-blocking network poller._  
_注意处于“IO 等待”状态的 goroutine 不消耗 OS 线程。它们停在非阻塞网络轮询器上。_

Explore using the **top**, **list**, **web** and **weblist** commands.  
使用 top、list、web 和 weblist 命令进行探索。

#### Comparing Profiles

Take a snapshot of the current heap profile. Then do the same for the cpu profile.  
拍摄当前堆配置文件的快照。然后对 cpu 配置文件执行相同操作。   

    $ curl -s http://localhost:4000/debug/pprof/heap > base.heap

After some time, take another snapshot:  
一段时间后，拍摄另一个快照：

    $ curl -s http://localhost:4000/debug/pprof/heap > current.heap

Now compare both snapshots against the binary and get into the pprof tool:  
现在将两个快照与二进制文件进行比较并进入 pprof 工具：

    $ go tool pprof -inuse_space -base base.heap current.heap

### Benchmark Profiling

Run the benchmarks and produce a cpu and memory profile.  
运行基准测试并生成 CPU 和内存配置文件。

	$ cd ./search
	
	$ go test -run none -bench . -benchtime 3s -benchmem -cpuprofile p.out
	$ go tool pprof p.out
	(pprof) web list match

	$ go test -run none -bench . -benchtime 3s -benchmem -memprofile p.out
	$ go tool pprof -inuse_space p.out
	(pprof) web list match

### Trace Profiles

#### Trace Web Application

Capture a trace file for a brief duration.  
在短时间内捕获跟踪文件。

	$ curl -s http://localhost:4000/debug/pprof/trace?seconds=2 > trace.out

Run the Go trace tool.

	$ go tool trace trace.out

Use the RSS Search test instead to create a trace.

	$ cd search
	$ go test -run none -bench . -benchtime 3s -trace trace.out
	$ go tool trace trace.out

## Expvar

Package expvar provides a standardized interface to public variables, such as operation counters in servers. It exposes these variables via HTTP at /debug/vars in JSON format.  
包 expvar 提供了公共变量的标准化接口，例如服务器中的操作计数器。它通过 HTTP 在 debugvars 中以 JSON 格式公开这些变量。 

### Adding New Variables

	import "expvar"

	// expvars is adding the goroutine counts to the variable set.
	func expvars() {

		// Add goroutine counts to the variable set.
		gr := expvar.NewInt("Goroutines")
		go func() {
			for _ = range time.Tick(time.Millisecond * 250) {
				gr.Set(int64(runtime.NumGoroutine()))
			}
		}()
	}

	// main is the entry point for the application.
	func main() {
		expvars()
		service.Run()
	}

### Expvarmon

TermUI based Go apps monitor using expvars variables (/debug/vars). Quickest way to monitor your Go app.  
基于 TermUI 的 Go 应用程序使用 expvars 变量（调试变量）进行监控。监控您的 Go 应用的最快捷方式。

	$ go get github.com/divan/expvarmon

Running expvarmon

	$ expvarmon -ports=":5000" -vars="requests,goroutines,mem:memstats.Alloc"
