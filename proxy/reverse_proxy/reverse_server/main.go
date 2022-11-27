package main

import (
   "bufio"
   "log"
   "net/http"
   "net/url"
)

var (
   proxyAddr = "http://127.0.0.1:2003"
   port      = "2002"
)

func main() {
   http.HandleFunc("/", handler)
   log.Println("Start serving on port " + port)
   err := http.ListenAndServe(":"+port, nil)
   if err != nil {
      log.Fatal(err)
   }
}

func handler(w http.ResponseWriter, r *http.Request) {
   // step1: 解析代理地址，并更改请求体的协议和主机
   proxy, err := url.Parse(proxyAddr)
   r.URL.Scheme = proxy.Scheme
   r.URL.Host = proxy.Host

   // step2: 请求下游
   transport := http.DefaultTransport
   resp, err := transport.RoundTrip(r)
   if err != nil {
      log.Print(err)
      return
   }

   // step3: 把下游请求内容返回给上游
   for k, v := range resp.Header {
      for _, v := range v {
         w.Header().Add(k, v)
      }
   }
   defer resp.Body.Close()
   bufio.NewReader(resp.Body).WriteTo(w)
}
