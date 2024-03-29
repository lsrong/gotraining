## Blocking Profiling

Testing and Tracing allows us to see blocking profiles.  
测试和跟踪允许我们查看阻塞配置文件。

## Running a Test Based Blocking Profile

We can get blocking profiles by running a test.

Generate a block profile from running the test.

	$ go test -blockprofile block.out

Run the pprof tool to view the blocking profile.

	$ go tool pprof block.out

Review the TestLatency function.

	$ list TestLatency

## Running a Trace

Once you have a test established you can use the **-trace trace.out** option with the go test tool.

Generate a trace from running the test.

	$ go test -trace trace.out

Run the trace tool to review the trace.

	$ go tool trace trace.out

## Links

No Extra links at this time.

## Code Review

[Blocking Trace](blocking_test.go)