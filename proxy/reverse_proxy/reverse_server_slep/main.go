// http 反向代理，相比于reverse_server包中的代理方式，httputil.NewSingleHostReverseProxy() 更简洁
package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
)

var addr = "127.0.0.1:2002"

func main() {
	rs1 := "http://127.0.0.1:2003/base"

	url1, err1 := url.Parse(rs1)
	if err1 != nil {
		log.Println(err1)
	}

	proxy := NewSingleHostReverseProxy(url1)
	log.Println("starting http server at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}

func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
		if _, ok := req.Header["User-Agent"]; !ok {
			// explicitly disable User-Agent so it's not set to default value
			req.Header.Set("User-Agent", "")
		}
	}

	// 内容改写
	modifyFunc := func(r *http.Response) error {
		if r.StatusCode != 200 {
			oldPayload, err := io.Copy(os.Stdin, r.Body)
			if err != nil {
				return err
			}
			newPayload := []byte("hello " + string(oldPayload))
			r.Body = io.NopCloser(bytes.NewBuffer(newPayload))
			r.ContentLength = int64(len(newPayload))
			r.Header.Set("Content-Length", fmt.Sprint(len(newPayload)))
		}
		return nil
	}
	return &httputil.ReverseProxy{Director: director, ModifyResponse: modifyFunc}
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}
