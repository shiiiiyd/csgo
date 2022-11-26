package main

import (
   "fmt"
   "io"
   "net"
   "net/http"
   "strings"
)

type Pxy struct {
}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
   fmt.Printf("Received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
   transport := http.DefaultTransport
   // step1: 浅拷贝对象，然后在新增属性数据
   outReq := new(http.Request)
   *outReq = *req
   if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
      if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
         clientIP = strings.Join(prior, ", ") + ", " + clientIP
      }
      outReq.Header.Set("X-Forwarded-For", clientIP)
   } else {
      rw.WriteHeader(http.StatusBadGateway)
      return
   }

   // step2: 请求下载
   res, err := transport.RoundTrip(outReq)
   if err != nil {
      rw.WriteHeader(http.StatusBadGateway)
      return
   }

   // step3: 将下游请求内容返回给上游
   for key, value := range res.Header {
      for _, v := range value {
         rw.Header().Add(key, v)
      }
   }
   rw.WriteHeader(res.StatusCode)
   io.Copy(rw, res.Body)
   res.Body.Close()

}

func main() {
   fmt.Println("Serve on :8080")
   http.Handle("/", &Pxy{})
   http.ListenAndServe("0.0.0.0:8080", nil)
}
