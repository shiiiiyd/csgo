[TOC]

### 1、Go test

Go 语言自带测试，测试代码和正式代码放在同一个目录。

> addition_test.go

```go
import "testing"

func TestAddition(t *testing.T) {
  if Addition(2, 3) == 5 {
    fmt.Println("2 + 3 = 5")
  }
}
```

**Go 自带的测试命令：**

```ba
# go test ./.. -v
go test addition_test.go -v
```

### 2、Go vet

代码静态检查，发现可能的 bug 或者可疑的构造

1. Print-format 错误，检查类型不匹配的print，例如：

```go
str := "hello world!"
fmt.Printf("%d\n", str)
```

2. Boolean 错误，检查一直为true、false或者冗余的表达式

```go
fmt.Println(i != 0 || i != 1)
```

3. Range 循环，比如下面代码主协程会先退出，goroutine 无法被执行

```go
words := []string{"foo","bar","baz"}
for _, word := range words{
  go func(){
    fmt.Println(word)
  }()
}
```

4. Unreachable 的代码，如 return 之后的代码
5. 变量自赋值，error检查滞后等

```go
res, err := http.Get("https://wwww.docker.io")
defer res.Body.Close()
if err != nil {
  log.Fatal(err)
}
```


