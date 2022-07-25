[toc]

2021年4月17

### 一、Goroutine

在Go语言中，每一个并发的执行单元叫作一个goroutine，当一个程序启动时，其主函数即在一个单独的goroutine中运行，我们叫它main goroutine。新的 goroutine会用go语句来创建。在语法上，go语句是一个普通的函数或方法调用前加上关键字go。

```go
f() // call f(); wait for it to return
go f() // create a new goroutine that calls f(); don't wait
```

**Thread vs. Groutine**

1. 创建时默认的stack大小

   - Jdk5 以后 java Thread stack 默认为 1 M

   - Groutine 的 stack 初始化大小为 2k

2. 和 KSE（Kernel Space Entity）的对应关系：

   - java Thread 时1:1

   - Groutine 是 M:N



### 二、fix 

**为了防止协程泄漏使用buffered channel**
