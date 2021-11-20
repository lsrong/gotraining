package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// rwMutex 读写锁
/**
// A RWMutex is a reader/writer mutual exclusion lock.
// The lock can be held by an arbitrary number of readers or a single writer.
// The zero value for a RWMutex is an unlocked mutex.
//
// A RWMutex must not be copied after first use.
//
// If a goroutine holds a RWMutex for reading and another goroutine might
// call Lock, no goroutine should expect to be able to acquire a read lock
// until the initial read lock is released. In particular, this prohibits
// recursive read locking. This is to ensure that the lock eventually becomes
// available; a blocked Lock call excludes new readers from acquiring the
// lock.
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // semaphore for writers to wait for completing readers
	readerSem   uint32 // semaphore for readers to wait for completing writers
	readerCount int32  // number of pending readers
	readerWait  int32  // number of departing readers
}
RWMutex是基于Mutex的，在Mutex的基础之上增加了读、写的信号量，并使用了类似引用计数的读锁数量
读锁与读锁兼容，读锁与写锁互斥，写锁与写锁互斥，只有在锁释放后才可以继续申请互斥的锁：
	可以同时申请多个读锁
	有读锁时申请写锁将阻塞，有写锁时申请读锁将阻塞
	只要有写锁，后续申请读锁和写锁都将阻塞

可以随机申请多个读锁，有未解除的读锁的时候，写锁会等待读锁的完成，有写锁的时候申请读锁会阻塞。
可以存在多个随机读锁，但是不予许存在读写锁并存，

Mutex和RWMutex都不关联goroutine，但RWMutex显然更适用于读多写少的场景。仅针对读的性能来说，RWMutex要高于Mutex，因为RWMutex的多个读可以并存。
*/
var rwMutex sync.RWMutex

// data is a slice that will be shared
var data []string

// readCount Number of reads occurring at given time
var readCount int64

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// Create a writer goroutine
	go func() {
		for i := 0; i < 10; i++ {
			writer(i)
		}
		wg.Done()
	}()

	// Create eight reader goroutines
	wg.Add(8)
	for i := 0; i < 8; i++ {
		go func(i int) {
			for {
				reader(i)
			}
		}(i)
		wg.Done()
	}

	// Wait for the write goroutine to finish
	wg.Wait()
	fmt.Println("Program competed!")
}

// writer adds a new string to the slice in random intervals
func writer(i int) {
	rwMutex.Lock()
	{
		rc := atomic.LoadInt64(&readCount)
		// 写入时间随机等待会比较大
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Printf("****> : Performing Wrire: RCount: %d \n", rc)
		data = append(data, fmt.Sprintf("String: %d", i))
	}
	rwMutex.Unlock()
}

// reader wake up and iterates over the data slice
func reader(id int) {
	rwMutex.RLock()
	{
		// data 为改变时候读取的数量 +1
		rc := atomic.AddInt64(&readCount, 1)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

		fmt.Printf("%d : Performing Read : Length[%d] RCount[%d] \n", id, len(data), rc)
		// 减一恢复
		atomic.AddInt64(&readCount, -1)
	}
	rwMutex.RUnlock()
}
