目录

[toc]

2021年4月13

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

   - java Thread 是1:1

   - Groutine 是 M:N

### 二、channels

channels则是 goroutine 之间的通信机制。一个 channel是一个通信机制，它可以让一个goroutine通过它给另一个goroutine发送值信息。每个 channel都有一个特殊的类型，也就是channels可发送数据的类型。一个可以发送 int 类型数据的 channel一般写为chan int。

使用内置的make函数，我们可以创建一个channel:

```go
ch := make(chan int) // ch has type 'chan int'
```

和map类似，channel也对应一个make创建的底层数据结构的引用。当我们复制一个channel或用于函数参数传递时，我们只是拷贝了一个channel引用，因此调用者和被调用者将引用同一个channel 对象。和其它的引用类型一样，channel的零值也是nil。

两个相同类型的channel可以使用==运算符比较。如果两个channel引用的是相同的对象，那么比较 的结果为真。一个channel也可以和nil进行比较。

#### （1）、closec操作

用于关闭channel，随后对基于该channel的任何发送操作都将导致 panic异常。对一个已经被close过的channel进行接收操作依然可以接受到之前已经成功发送的数 据;如果channel中已经没有数据的话将产生一个零值的数据（对应类型的零值）。使用内置的close函数关闭一个channel：

```go
close(ch)
```

#### （2）、bufered channel(带缓存的channel)

使用make函数创建的是一个无缓存的channel，但是我们也可以指定第二个整型参数， 对应channel的容量。如果channel的容量大于零，那么该channel就是带缓存的channel。例如：

```go
ch = make(chan int) // unbuffered channel
ch = make(chan int, 0) // unbuffered channel
# 带缓存的channel
ch = make(chan int, 3) // buffered channel with capacity 3
```

### 三、fix 错误

**1、为了防止协程泄漏使用buffered channel（带缓存的channel）**

```go
// 防止协程泄漏，使用buffered channel
package channel_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", id)
}

func BufferedChannel() string {
	numOfRunner := 10
	//ch := make(chan string)  // 会导致协程泄漏
	ch := make(chan string, numOfRunner) // buffered channel
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestBufferChannel(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(BufferedChannel())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}

```

