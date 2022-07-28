目录

[toc]

2021年4月22

### http的路由规则

- URL分为两种，末尾时 / 表示一个子树，后面可以跟其他子路径；末尾是/，表示一个叶子，固定的路径以 / 结尾的URL可以匹配它的任何子路径，比如 /images/ 会匹配 /images/cute-cat.jpg.
- 采用最长匹配原则，如果有多个匹配，一定采用匹配路径最长的那个进行处理。
- 如果没有找到任何匹配项，返回404错误。

```go
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
```

在浏览中输入：`localhost:8080/time/1` 返回具体的时间。输入`localhost:8080/1`，界面返回信息为 “Hello，世界”。因此采用了最长匹配原则。



