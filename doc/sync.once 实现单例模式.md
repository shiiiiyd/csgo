2021年4月27

### 一、Java中单例模式（懒汉式，线程安全）

```java
public class Singleton {
  private static Singleton INSTANCE = null;
  private Singleton(){}
  
  public static Singleton getIntance(){
    if(INSTANCE == null) {
      syncchronized(Singleton.class){
        if(INSTANCE == null) {
          INSTANCE = new Singleton();
        }
      }
    }
    return INSTANCE;
  }
}
```

### 二、Go 中单例模式的实现

sync.Once 中的 Do函数确保方法只调用一次，因此可以通过方法中创建实例，然后使用Do函数调用该方法，使其只执行一次。 

```go
// 使用 sync.Once 中的Do函数创建单例模式
package sync_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct{}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton {
  // 只执行一次func中的代码
	once.Do(func() {
		fmt.Println("Create Obj")
    // 创建一个实例
		singleInstance = new(Singleton)
	})
	return singleInstance
}

// 测试
func TestSingletonByOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

=== 输出结果
=== RUN   TestSingletonByOnce
Create Obj
1028bbef8
1028bbef8
1028bbef8
1028bbef8
1028bbef8
1028bbef8
1028bbef8
1028bbef8
1028bbef8
1028bbef8
--- PASS: TestSingletonByOnce (0.00s)
PASS
ok      csgo/goroutine/sync     0.125s

```

