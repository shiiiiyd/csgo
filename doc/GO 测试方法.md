[TOC]

### Go test

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

### Go vet

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



### 接口

注意事项：

- Interface 是有可能为nil的，所以针对 interface 的 使用一定需要预先判空，否则会引起程序crash（nil panic）
- Struct 初始化意味着空间分配，对 struct 的引用不会出现空指针

### 反射

### Json编解码

- Unmarshal: 从 string 转换至 struct
- Marshal: 从 struct 转换至 string
- 

### Defer

Defer 是在函数退出后执行，如果有多个defer，defer执行的顺序是先出现的后执行（类似于将defer放入在栈中）。

注意Defer是函数退出后执行，而不是程序退出后执行。

**并发和并行：** 并发是一件事的执行过程有多个线程执行，但是每个线程所拥有的时间间隔上不能有其他线程，同一时间只能有一个线程执行。并行是同一时间有多个线程执行。



### goroutine



### chanl



### mutex



### 内存管理

默认配置下，P的数量=CPU个数，可通过GOMAXPROCS修改，上限是256。

M数量是内核线程数量，M大于P，上限是10000。

P的数量在调度器初始化的procresize中控制。

当调度器进行调度，唤醒P的时候，会通过met获取idle m，如果获取不到，则创建新M也就是内核线程。

这样做的目的是，M可能陷入系统调用，而系统调用可能是阻塞的，比如磁盘读取，这个时候CPU是空闲的，创建新的M与P关联，可以让更多的G被调度，充分利用了CPU。

> 对内存管理

![image-20220721200025264](/Users/shiyd/Library/Application Support/typora-user-images/image-20220721200025264.png)

> **TCMalloc**

- page:内存页，一块8K 大小的内存空间。Go 与操作系统之间的内存申请和释放，都是以page 为单位的。
- span: 内存块，一个或多个连续的page 组成一个span•sizeclass:空间规格，每个span 都带有一个sizeclass，标记着该span 中的page 应该如何使用。
- object : 对象，用来存储一个变量数据内存空间，一个span 在初始化时，会被切割成一堆等大的object ；假设object 的大小是16B ，span 大小是8K ，那么就会把span 中的page 就会被初始化8K / 16B = 512 个object 。所谓内存分配，就是分配一个object 出去。

- 对象大小定义
  - 小对象大小：0~256KB•中对象大小：256KB~1MB。
  - 大对象大小：>1MB•小对象的分配流程。
  - ThreadCache-> CentralCache-> HeapPage，大部分时候，ThreadCache缓存都是足够的，不需要去访问CentralCache和HeapPage，无系统调用配合无锁分配，分配效率是非常高的
- 中对象分配流程
  - 直接在PageHeap中选择适当的大小即可，128 Page的Span所保存的最大内存就是1MB。
  - 大对象分配流程•从large span set选择合适数量的页面组成span，用来存储数据。

> **ThreadCacheMalloc 概述**

![image-20220721231931525](/Users/shiyd/Library/Application Support/typora-user-images/image-20220721231931525.png)

### Go 语言内存分配

![image-20220721232411515](/Users/shiyd/Library/Application Support/typora-user-images/image-20220721232411515.png)

- mcache：小对象的内存分配直接走
  - size class从1到66，每个class两个span。
  - Span大小是8KB，按spanclass大小切分。
- mcentral
  - Span内的所有内存块都被占用时，没有剩余空间继续分配对象，mcache会向mcentral申请1个span，mcache拿到span后继续分配对象。
  - 当mcentral向mcache提供span时，如果没有符合条件的span，mcentral会向mheap申请span•mheap。
  - 当mheap没有足够的内存时，mheap会向OS申请内存
- Mheap把Span组织成了树结构，而不是链表
  - 然后把Span分配到heapArena进行管理，它包含地址映射和span是否包含指针等位图
  - 为了更高效的分配、回收和再利用内存。

> **内存回收**

- 引用计数（Python，PHP，Swift）
  - 对每一个对象维护一个引用计数，当引用该对象的对象被销毁的时候，引用计数减1，当引用计数为0的时候，回收该对象•
  - 优点：对象可以很快的被回收，不会出现内存耗尽或达到某个阀值时才回收
  - 缺点：不能很好的处理循环引用，而且实时维护引用计数，有也一定的代价
- 标记-清除（Golang）
  - 从根变量开始遍历所有引用的对象，引用的对象标记为"被引用"，没有被标记的进行回收。
  - 优点：解决引用计数的缺点。
  - 缺点：需要STW（stop the word），即要暂停程序运行。
- 分代收集（Java）
  - 按照生命周期进行划分不同的代空间，生命周期长的放入老年代，短的放入新生代，新生代的回收频率高于老年代的频率。

> **mspan**

- allocBits
  - 记录了每块内存分配的情况
- gcmarkBits
  - 记录了每块内存的引用情况，标记阶段对每块内存进行标记，有对象引用的内存标记为1，没有的标记为0

![image-20220721234432859](/Users/shiyd/Library/Application Support/typora-user-images/image-20220721234432859.png)

- 这两个位图的数据结构是完全一致的，标记结束则进行内存回收，回收的时候，将allocBits指向gcmarkBits，标记过的则存在，未进行标记的则进行回收。

![image-20220721234534818](/Users/shiyd/Library/Application Support/typora-user-images/image-20220721234534818.png)

> **GC 工作流程**

Golang GC的大部分处理是和用户代码并行的

- Mark：
  - Mark Prepare: 初始化GC任务，包括开启写屏障(write barrier)和辅助GC(mutator assist)，统计root对象的任务数量等。这个过程需要STW。
  - GC Drains: 扫描所有root对象，包括全局指针和goroutine(G)栈上的指针（扫描对应G栈时需停止该G)，将其加入标记队列(灰色队列)，并循环处理灰色队列的对象，直到灰色队列为空。该过程后台并行执行。
- Mark Termination：完成标记工作，重新扫描(re-scan)全局指针和栈。因为Mark和用户程序是并行的，所以在Mark过程中可能会有新的对象分配和指针赋值，这个时候就需要通过写屏障（write barrier）记录下来，re-scan 再检查一下，这个过程也是会STW的。
- Sweep：按照标记结果回收所有的白色对象，该过程后台并行执行•Sweep Termination：对未清扫的span进行清扫, 只有上一轮的GC的清扫工作完成才可以开始新一轮的GC。

![image-20220721233259564](/Users/shiyd/Library/Application Support/typora-user-images/image-20220721233259564.png)

> **三色标记**

- GC 开始时，认为所有object 都是白色，即垃圾。
- 从root 区开始遍历，被触达的object 置成灰色。
- 遍历所有灰色object，将他们内部的引用变量置成灰色，自身置成黑色•循环第3 步，直到没有灰色object 了，只剩下了黑白两种，白色的都是垃圾。
- 对于黑色object，如果在标记期间发生了写操作，写屏障会在真正赋值前将新对象标记为灰色
- 标记过程中，mallocgc新分配的object，会先被标记成黑色再返回。

![image-20220721233435662](/Users/shiyd/Library/Application Support/typora-user-images/image-20220721233435662.png)

> **垃圾回收机制**

- 内存分配量达到阀值触发GC。

  每次内存分配时都会检查当前内存分配量是否已达到阀值，如果达到阀值则立即启动GC。

  - 阀值= 上次GC内存分配量* 内存增长率。
  - 内存增长率由环境变量GOGC控制，默认为100，即每当内存扩大一倍时启动GC。

- 定期触发GC

  默认情况下，最长2分钟触发一次GC，这个间隔在src/runtime/proc.go:forcegcperiod变量中被声明

- 手动触发

  程序代码中也可以使用runtime.GC()来手动触发GC。这主要用于GC性能测试和统计。