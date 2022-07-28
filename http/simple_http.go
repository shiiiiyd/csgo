package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 第一个参数时访问目录，第二个参数时匿名函数，实现handler逻辑处理控制，类似controller方法
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, 世界！")
	})

	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	http.ListenAndServe(":8080", nil)
}
