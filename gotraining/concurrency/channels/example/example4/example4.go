package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// This sample program demonstrates how to use a buffered
// channel to receive results from other goroutines in a guaranteed way.

// 此示例程序演示了如何使用缓冲通道以有保证的方式接收来自其他 goroutine 的结果。

type result struct {
	id  int
	op  string
	err error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	const routines = 10
	const inserts = routines * 2
	// 缓冲通道接受任何插入的结果
	ch := make(chan result, inserts)

	waitInserts := inserts

	for i := 0; i < routines; i++ {

		go func(id int) {
			ch <- insertUser(id)

			// 由于缓冲通道，我们不需要等待开始第二次插入。第一次发送将立即发生。
			ch <- insertTrans(id)
		}(i)
	}

	// Process the insert results as they complete.
	// 在插入结果完成时对其进行处理。
	for waitInserts > 0 {
		r := <-ch
		log.Printf("N: %d, ID: %d, OP:%s ERR: %v", waitInserts, r.id, r.op, r.err)
		waitInserts--
	}
	log.Println("Inserts Complete")
}

// insertUser 模拟插入User的操作。
func insertUser(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf("insert USERS value (%d)", id),
	}

	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into USER table ", id)
	}

	return r
}

// insertTrans 模拟插入Trans操作
func insertTrans(id int) result {
	r := result{
		id: id,
		op: fmt.Sprintf("insert TRANS value (%d)", id),
	}

	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("Unable to insert %d into TRANS table ", id)
	}

	return r
}
