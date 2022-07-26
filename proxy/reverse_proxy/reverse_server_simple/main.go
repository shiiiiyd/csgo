// http 反向代理，相比于reverse_server包中的代理方式，httputil.NewSingleHostReverseProxy() 更简洁
package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var addr = "127.0.0.1:2002"

func main() {
	rs1 := "http://127.0.0.1:2003/base"

	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}

	proxy := httputil.NewSingleHostReverseProxy(url1)
	log.Println("starting http server at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
