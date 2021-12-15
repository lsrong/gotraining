package queue

import (
	"errors"
	"github.com/learning_golang/algorithms/data/list"
)

// Queue 队列数据结构，特点为先进先出（FIFO，First In First Out）的方式存储和检索元素。
// 常用链表来实现队列的功能。
type Queue struct {
	list.List
}

// NewQueue 初始化队列。
func NewQueue() *Queue {
	var l list.List

	return &Queue{l}
}

// Enqueue 入队操作，将新元素添加到队列中
func (q *Queue) Enqueue(data interface{}) int {
	// 插入到最后一位，作为入队
	q.InsertAfter(data, q.Tail())

	return q.Len()
}

// Dequeue 出队操作，从队列的头部删除一个元素
func (q *Queue) Dequeue() (interface{}, error) {
	// 弹出首位元素，作为出队。
	node, err := q.RemoveAfter(nil)
	if err != nil {
		return nil, errors.New("queue is empty")
	}

	return node.Data(), nil
}

// Peek 返回队列的头部元素
func (q *Queue) Peek() interface{} {
	if head := q.List.Head(); head != nil {
		return head.Data()
	}

	return nil
}

func (q *Queue) Len() int {
	return q.List.Len()
}
