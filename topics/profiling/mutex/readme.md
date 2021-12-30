## Mutex Profiling

Testing and Tracing allows us to see mutex profiles.  
测试和跟踪允许我们查看互斥锁配置文件。

## Running a Test Based Mutex Profile

We can get mutex profiles by running a test.

Generate a mutex profile from running the test.

	$ go test -mutexprofile mutex.out

Run the pprof tool to view the mutex profile.

	$ go tool pprof mutex.out

Review the TestMutexProfile function.

	$ list TestProfileMutex

## Links

[Mutex profile](https://rakyll.org/mutexprofile) - JBD

## Code Review

[Mutex Trace](mutex_test.go)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).
