package main

import (
	"fmt"
	"github.com/shiiiiyd/csgo/go_learning/src/unpack/unpack"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err: %v\n", err.Error())
		return
	}
	unpack.Encode(conn, "hello world!!")
}
