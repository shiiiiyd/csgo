// 对象缓存池
package objcache

import (
    "fmt"
    "runtime"
    "sync"
    "testing"
)

func TestSyncPool(t *testing.T) {
    pool := &sync.Pool{
        // interface{}
        New: func() interface{} {
            fmt.Println("Create a new object.")
            return 100
        },
    }

    // .(int) 断言，由于返回的是 interface{}，所以无法确定类型，需要使用类型断言指定类型，
    // 否则会出现错误
    v := pool.Get().(int)
    fmt.Println(v)
    pool.Put(3)

    //GC 会清除sync.pool中缓存的对象，节省内存空间，但也会消耗时间
    runtime.GC()
    v1, _ := pool.Get().(int)
    fmt.Println(v1)
}

func TestSyncPoolInMultiGoroutine(t *testing.T) {
    pool := &sync.Pool{
        New: func() interface{} {
            fmt.Println("Crate a new object.")
            return 10
        },
    }

    pool.Put(100)
    pool.Put(100)
    pool.Put(100)

    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(id int) {
            fmt.Println(pool.Get())
            wg.Done()
        }(i)
    }
    wg.Wait()
}
