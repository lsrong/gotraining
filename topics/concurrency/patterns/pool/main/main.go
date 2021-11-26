package main

// This sample program demonstrates how to use the pool package
// to share a simulated set of database connections.

// 这个示例程序演示了如何使用资源池包来共享一组模拟的数据库连接。

import (
	"fmt"
	"io"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/learning_golang/topics/concurrency/patterns/pool"
)

// dbConnection 模拟数据库连接句柄
type dbConnection struct {
	id int32
}

// Close 模拟关闭操作
func (dbConn *dbConnection) Close() error {
	log.Println("close connection ", dbConn.id)
	return nil
}

var idCounter int32

// connect 模拟数据库的连接函数，返回资源连接对象
func connect() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)

	return &dbConnection{id: id}, nil
}

// performQueries 执行一次数据查询
func performQueries(query int, p *pool.Pool) {
	// 获得一个资源对象
	conn, err := p.Acquired()
	if err != nil {
		log.Println(err)
		return
	}

	// 放回资源池
	defer p.Release(conn)

	// 执行操作
	time.Sleep(time.Second * 3)
	log.Printf("Query:QID[%d] CID[%d] \n", query, conn.(*dbConnection).id)
}

func main() {
	const maxGoroutines = 25
	const numPools uint = 2

	var wg sync.WaitGroup

	// 初始化资源池
	p, err := pool.New(numPools, connect)
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(maxGoroutines)

	// 模拟执行一次数据库操作
	for i := 0; i < maxGoroutines; i++ {
		go func(query int) {
			performQueries(query, p)
			wg.Done()
		}(i)
	}

	wg.Wait()
	// 关闭资源池
	err = p.Close()
	fmt.Println("Resource pool is closed")
}
