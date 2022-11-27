package main

import (
   "fmt"
   "io"
   "log"
   "net/http"
   "os"
   "os/signal"
   "syscall"
   "time"
)

type ReverseServer struct {
   Addr string
}

func (r ReverseServer) Run() {
   log.Println("Starting httpserver at " + r.Addr)
   mux := http.NewServeMux()
   mux.HandleFunc("/", r.HelloHandler)
   mux.HandleFunc("/base/error", r.ErrorHandler)
   server := &http.Server{
      Addr:         r.Addr,
      WriteTimeout: time.Second * 3,
      Handler:      mux,
   }
   go func() {
      log.Fatal(server.ListenAndServe())
   }()
}

func (r ReverseServer) HelloHandler(w http.ResponseWriter, req *http.Request) {
   upath := fmt.Sprintf("http://%s%s\n", r.Addr, req.URL.Path)
   io.WriteString(w, upath)
}

func (r ReverseServer) ErrorHandler(w http.ResponseWriter, req *http.Request) {
   upath := "error handler"
   w.WriteHeader(500)
   io.WriteString(w, upath)
}

func main() {
   rs1 := &ReverseServer{Addr: "127.0.0.1:2003"}
   rs1.Run()
   rs2 := &ReverseServer{Addr: "217.0.0.1:2004"}
   rs2.Run()

   // 监听关闭信号
   quit := make(chan os.Signal)
   signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
   <-quit
}
