package main

import (
   "fmt"
   "io"
   "net"
   "net/http"
   "os"
   "time"
)

func main() {
   // 创建连接池
   transport := &http.Transport{
      DialContext: (&net.Dialer{
         Timeout:   30 * time.Second, // 连接超时
         KeepAlive: 30 * time.Second,
      }).DialContext,
      MaxIdleConns:          100,
      IdleConnTimeout:       90 * time.Second, // 空闲超时时间
      TLSHandshakeTimeout:   10 * time.Second, // tls握手超时时间
      ExpectContinueTimeout: 1 * time.Second,
   }
   // 创建客户端
   client := &http.Client{
      Timeout:   time.Second * 30,
      Transport: transport,
   }
   // 请求数据
   resp, err := client.Get("http://127.0.0.1:1210/bye")
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()
   // 读取内容
   //bds, err := ioutil.ReadAll(resp.Body)
   // 使用 io.Copy 代替 ioutil.ReadAll 方法
   bds, err := io.Copy(os.Stdout, resp.Body)
   if err != nil {
      panic(err)
   }
   fmt.Println(string(bds))
}
