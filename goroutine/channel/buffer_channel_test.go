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
