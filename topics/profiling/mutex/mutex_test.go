package mutex

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

// 如何分析互斥锁的示例程序。

var rwMutex sync.RWMutex
var data []string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestProfileMutex(t *testing.T) {
	t.Log("Starting Test")
	var wg sync.WaitGroup
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			writer()
			wg.Done()
		}()

		go func() {
			reader()
			wg.Done()
		}()
	}

	wg.Wait()
	t.Log("Test Complete!")
}

// writer 模拟写数据
func writer() {
	for i := 0; i < 10; i++ {
		rwMutex.Lock()
		{
			data = append(data, fmt.Sprintf("node %d", i))
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}
		rwMutex.Unlock()
	}
}

func reader() {
	for i := 0; i < 10; i++ {
		rwMutex.RLock()
		{
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		}
		rwMutex.RUnlock()
	}
}
