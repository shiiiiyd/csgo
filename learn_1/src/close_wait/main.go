// close wait.
package main

import (
   "fmt"
   "net"
)

func main() {
   // 1. 监听端口
   listener, err := net.Listen("tcp", "0.0.0.0:9090")
   if err != nil {
      fmt.Println(err)
   }
   // 2. 建立websocket连接
   for {
      conn, err := listener.Accept()
      if err != nil {
         fmt.Println(err)
      }
      // 3.创建处理协程
      go func(conn net.Conn) {
         defer conn.Close()
         for {
            var buf [128]byte
            n, err := conn.Read(buf[:])
            if err != nil {
               fmt.Printf("read from connect failed,err:%v\n", err)
               break
            }
            str := string(buf[:n])
            fmt.Printf("receive from client, data:%v\n", str)
         }
      }(conn)
   }
}
