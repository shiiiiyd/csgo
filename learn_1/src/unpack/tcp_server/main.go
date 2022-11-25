package main

import (
	"fmt"
	"github.com/shiiiiyd/csgo/go_learning/src/unpack/unpack"
	"net"
)

func main() {
	//	simple tcp server
	//	1.listen ip+port
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}

	// 2. accept client request
	//	3.create goroutine each request
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err:%v\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		str := string(bt)
		fmt.Printf("receive from client, data: %v\n", str)
	}
}
