package pool

// Package pool manages a user defined set of resources.
// 管理用户定义的资源池。

import (
	"errors"
	"io"
	"log"
	"sync"
)

// Pool 资源管理。
type Pool struct {
	mu        sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// ErrPoolClosed 资源已关闭错误
var ErrPoolClosed = errors.New("Pool has been closed ")

// New 初始化资源池
func New(size uint, f func() (io.Closer, error)) (*Pool, error) {
	if size == 0 {
		return nil, errors.New("Size isn't zero ")
	}
	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   f,
	}, nil
}

// Acquired 从资源池中获取一个资源连接，如果没有空闲资源就调用工厂函数生成新资源。
func (p *Pool) Acquired() (io.Closer, error) {
	select {
	case conn, ok := <-p.resources:
		// 判断资源池是否已经关闭
		if !ok {
			return nil, ErrPoolClosed
		}
		return conn, nil
	default:
		return p.factory()
	}
}

// Release 已经使用的资源连接放入资源池，如果到达容量则关闭连接。
func (p *Pool) Release(r io.Closer) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resources <- r:
		log.Println("Release In Queue")
	default:
		log.Println("Release Closing")
		r.Close()
	}
}

// Close 关闭资源池.关闭资源队列（channel）,关闭空闲连接。
func (p *Pool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	// 已经关闭
	if p.closed {
		return ErrPoolClosed
	}

	// 关闭资源channel
	close(p.resources)

	// 遍历资源channels,关闭空闲的资源资源连接
	for r := range p.resources {
		r.Close()
	}

	return nil
}
