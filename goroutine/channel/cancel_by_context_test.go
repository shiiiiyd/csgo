// 通过 context 取消任务，连同子任务一起取消
package channel_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done(): // 接受取消通知 ctx.Done()
		return true
	default:
		return false
	}
}

func TestCancelByContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)

}
