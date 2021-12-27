package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

// 实现一个简单的 Web 服务的示例程序并探索如何使用 GODEBUG 变量

var leak bool

func main() {
	http.HandleFunc("/json", sendJSON)
	// 当有第二个参数为leak的时候 开启goroutine泄露，
	if len(os.Args) == 2 && os.Args[1] == "leak" {
		leak = true
	}
	log.Println("Server listening on http://0.0.0.0:4000")

	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}

// sendJSON  响应一个json文本的数据。
func sendJSON(w http.ResponseWriter, r *http.Request) {
	if leak {
		if rand.Intn(100) == 5 {
			go func() {
				for {
					time.Sleep(time.Millisecond)
				}
			}()
		}
	}
	u := struct {
		Name  string
		Email string
	}{
		"LiuShengRong", "lsrong0414@gmail.com",
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(&u)
}

// 构建：
// $ go build .
// $ GOMAXPROCS=1 GODEBUG=schedtrace=1000 ./godebug leak

// 看到追踪输出为：
// SCHED 0ms: gomaxprocs=1 idleprocs=0 threads=3 spinningthreads=0 idlethreads=0 runqueue=0 [0]
//2021/12/27 14:58:26 Server listening on http://0.0.0.0:4000
//SCHED 1006ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 idlethreads=1 runqueue=0 [0]
//SCHED 2006ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 idlethreads=1 runqueue=0 [0]
//SCHED 3009ms: gomaxprocs=1 idleprocs=1 threads=5 spinningthreads=0 idlethreads=1 runqueue=0 [0]
// 上面输出的说明：
// sched: 启动到现在输出行的运行时间，
// gomaxprocs: 当前运行的cpu核心数
// idleprocs: 空闲核心数
// threads：前正在运行的OS线程数
// spinningthreads: 自旋状态的 OS 线程数量。
// idlethreads：空闲的线程数量。
// runqueue：全局队列中的 Goroutine 数量，而后面的 [0] 则分别代表这 1 个 P 的本地队列正在运行的 Goroutine 数量。

// 模拟压测请求,让服务出现goroutine泄露情况
// hey -m POST -c 8 -n 1000000 "http://localhost:4000/json"

// 再次打印
//SCHED 316844ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 idlethreads=1 runqueue=866 [179]
//SCHED 317849ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 idlethreads=1 runqueue=1170 [48]
//SCHED 318858ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 idlethreads=1 runqueue=1133 [217]
//SCHED 319859ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 idlethreads=1 runqueue=1292 [253]
//SCHED 320864ms: gomaxprocs=1 idleprocs=0 threads=5 spinningthreads=0 idlethreads=1 runqueue=1000 [125]

// GC Trace
// $ go build .
// $ GODEBUG=gctrace=1 ./godebug
// $ hey -m POST -c 8 -n 1000000 "http://localhost:4000/sendjson"

/**
	查看GCTrace信息：

    gc 318 @36.750s 0%: 0.022+0.27+0.040 ms clock, 0.13+0.60/0.43/0.031+0.24 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
    gc 319 @36.779s 0%: 0.019+0.24+0.035 ms clock, 0.15+0.43/0.26/0+0.28 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
    gc 320 @36.806s 0%: 0.023+0.34+0.035 ms clock, 0.18+0.63/0.49/0.014+0.28 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
    gc 321 @36.834s 0%: 0.026+0.20+0.044 ms clock, 0.18+0.50/0.34/0.001+0.31 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
    gc 322 @36.860s 0%: 0.022+0.29+0.055 ms clock, 0.13+0.54/0.47/0+0.33 ms cpu, 4->4->0 MB, 5 MB goal, 8 P

    gc 318      : First GC run since program started.
    @36.750s    : Nine milliseconds since the program started.
    0%          : One percent of the programs time has been spent in GC.

    // wall-clock
    0.022ms     : **STW** Sweep termination - Wait for all Ps to reach a GC safe-point.
    0.27ms      : Mark/Scan
    0.040ms     : **STW** Mark termination - Drain any remaining work and perform housekeeping.

    // CPU time
    0.13ms      : **STW** Sweep termination - Wait for all Ps to reach a GC safe-point.
    0.60ms      : Mark/Scan - Assist Time (GC performed in line with allocation)
    0.43ms      : Mark/Scan - Background GC time
    0.031ms     : Mark/Scan - Idle GC time
    0.24ms      : **STW** Mark termination - Drain any remaining work and perform housekeeping.

    4MB         : Heap size at GC start
    4MB         : Heap size at GC end
    0MB         : Live Heap
    5MB         : Goal heap size
    8P          : Number of logical processors
*/
