package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

// 模拟一次线程io-bound阻塞的场景
// IO-Bound：这是导致线程进入等待状态的工作。这项工作包括通过网络请求访问资源或对操作系统进行系统调用。
// 需要访问数据库或者有io等待的线程将是 IO-Bound。我将包括同步事件（互斥体、原子），这会导致线程作为此类别的一部分等待。

func main() {
	docs := generateDocs(1e3)
	fmt.Printf("Sequential found: %d \n", find("Go", docs))
	fmt.Printf("Concurrent found: %d \n", findConcurrent(runtime.NumCPU(), "Go", docs))

}

// generateDocs 生成批量io任务
func generateDocs(totalDocs int) []string {
	docs := make([]string, totalDocs)
	for i := 0; i < totalDocs; i++ {
		docs[i] = "test.xml"
	}

	return docs
}

// find 顺序同步读取文件io
func find(search string, docs []string) int {
	var found int
	for _, doc := range docs {
		items, err := Read(doc)
		if err != nil {
			continue
		}
		for _, item := range items {
			if strings.Contains(item.Description, search) {
				found++
			}

		}
	}

	return found
}

// findConcurrent 并发同步读取文件io
func findConcurrent(goroutines int, search string, docs []string) int {
	var found int64
	var wg sync.WaitGroup
	wg.Add(goroutines)
	docCh := make(chan string, len(docs))
	for i := 0; i < goroutines; i++ {
		go func() {
			var gFound int64
			for doc := range docCh {
				items, err := Read(doc)
				if err != nil {
					continue
				}
				for _, item := range items {
					if strings.Contains(item.Description, search) {
						gFound++
					}
				}
			}
			atomic.AddInt64(&found, gFound)
			wg.Done()
		}()
	}

	for _, doc := range docs {
		docCh <- doc
	}
	close(docCh)

	wg.Wait()

	return int(found)
}
