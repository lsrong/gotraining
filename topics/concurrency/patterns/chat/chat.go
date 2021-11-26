package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// message 消息主体
type message struct {
	data string
	conn net.Conn
}

// client 进入房间服务的客户端
type client struct {
	name   string
	conn   net.Conn
	room   *RoomServer
	reader *bufio.Reader
	writer *bufio.Writer

	wg sync.WaitGroup
}

// newClient 创建客户端连接实体, 启动读取信息的goroutine
func newClient(name string, conn net.Conn, room *RoomServer) *client {
	c := client{
		name:   name,
		conn:   conn,
		room:   room,
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
	}
	c.wg.Add(1)

	go c.read()

	return &c
}

// read 读取客户端连接信息，默认读取每行信息数据并发送给其他客户端
func (c *client) read() {
	for {
		// read one line conn message
		data, err := c.reader.ReadString('\n')

		// if err done goroutine
		if err != nil {
			if err == io.EOF {
				log.Printf("EOF: Client[%s] leaving chat \n", c.name)
			}
			c.wg.Done()
			return
		}

		// send conn one line message to msgCh
		c.room.msgCh <- message{
			data: data,
			conn: c.conn,
		}
	}
}

// write 向客户端写入消息.
func (c *client) write(m message) {
	data := fmt.Sprintf("%s: %s", c.name, m.data)
	log.Print(data)
	c.writer.WriteString(data)
	c.writer.Flush()
}

// drop 主动断开客户端连接
func (c *client) drop() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("client conn.Close err: %v", err)
	}
	c.wg.Wait()

	return nil
}

// RoomServer 聊天室主体，开启聊天服务，管理客户端连接，发送消息通道。
type RoomServer struct {
	listener net.Listener
	clients  []*client
	joining  chan net.Conn
	msgCh    chan message
	shutdown chan struct{}

	wg sync.WaitGroup
}

// New 创建并启动聊天室服务.
func New() *RoomServer {
	room := &RoomServer{
		joining:  make(chan net.Conn),
		msgCh:    make(chan message),
		shutdown: make(chan struct{}),
	}
	room.start()

	return room
}

// Close 关闭聊天室服务
func (r *RoomServer) Close() error {
	// close listener
	if err := r.listener.Close(); err != nil {
		return err
	}

	// close handle
	close(r.shutdown)

	r.wg.Wait()

	// drop clients
	for _, c := range r.clients {
		if err := c.drop(); err != nil {
			log.Printf("Client[%s] drop err: %v", c.name, err)
		}
	}

	return nil
}

// start 开启聊天服务，两个goroutines(handle, server), handle负责处理对应的事件信号，server负责tcp服务
func (r *RoomServer) start() {
	r.wg.Add(2)

	// handler goroutine
	go func() {
		for {
			select {
			case conn := <-r.joining:
				// revive a client
				r.join(conn)

			case msg := <-r.msgCh:
				// send message
				r.sendMessage(msg)

			case <-r.shutdown:
				// shutdown handle
				r.wg.Done()
				return
			}
		}
	}()

	// server goroutine
	go func() {
		var err error
		r.listener, err = net.Listen("tcp", ":1000")
		if err != nil {
			log.Fatalln(err)
		}

		log.Println("Chat room started: 1000")
		for {
			conn, err := r.listener.Accept()

			if err != nil {
				log.Println("Chat room shutting down")
				r.wg.Done()
				return
			}

			r.joining <- conn
		}
	}()

}

// join 客户端连接加入到聊天室服务.
func (r *RoomServer) join(conn net.Conn) {
	name := fmt.Sprintf("conn %d", len(r.clients))
	c := newClient(name, conn, r)

	r.clients = append(r.clients, c)
}

// sendMessage 发送消息.
func (r *RoomServer) sendMessage(m message) {
	for _, c := range r.clients {
		if c.conn != m.conn {
			c.write(m)
		}
	}
}
