package http

import (
	"fmt"
	"net"
)

func Start() {
	listen, err := net.Listen("tcp", "0.0.0.0:10000")
	if err != nil {
		fmt.Println("tcp listen failure!")
		return
	}

	for {
		connect, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept failure!")
			continue
		}
		// 处理http连接
		go Progress(connect)
	}
}
func Progress(connect net.Conn) {
	defer connect.Close()
	for {
		var buffer []byte
		n, err := connect.Read(buffer[:])
		if err != nil {
			fmt.Println("read connect failure")
			break
		}
		fmt.Printf("recv from connect:%s", string(buffer[:n]))
	}
}
