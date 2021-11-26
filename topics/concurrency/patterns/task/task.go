package task

// Package task provides a pool of goroutines to perform tasks.
// 提供了一个 goroutine 池来执行任务。

import "sync"

// Worker 任务执行接口
type Worker interface {
	Work()
}

// TManager 任务协程池管理对象，处理提交的不同 Worker
type TManager struct {
	workers chan Worker
	wg      sync.WaitGroup
}

// New 初始化工作任务协程池，返回指正对象。
func New(grs int) *TManager {
	tm := TManager{
		workers: make(chan Worker),
	}
	tm.wg.Add(grs)
	for i := 0; i < grs; i++ {
		// 启动一个协程监听工作channels, 执行提交的工作任务。
		go func() {
			for w := range tm.workers {
				w.Work()
			}

			tm.wg.Done()
		}()
	}
	return &tm
}

// Do 提交一次工作任务，加入到工作任务队列提供给协程池处理。
func (t *TManager) Do(w Worker) {
	t.workers <- w
}

// Shutdown 关闭工作池。
func (t *TManager) Shutdown() {
	close(t.workers)

	t.wg.Wait()
}
