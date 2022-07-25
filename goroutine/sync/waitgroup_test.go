package sync_test

import (
	"sync"
	"testing"
)

// 使用waitgroup 依次执行goroutine，之前使用的是sleep函数
func TestWaitGroup(t *testing.T) {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mutex.Unlock()
			}()
			mutex.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter = %d\n", counter)
}
