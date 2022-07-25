package goroutine_test

import (
	"fmt"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {

		// 这里的 i 是有竞争关系的，直接从 使用 for循环的 i
		// go func() {
		// 	fmt.Println(i)
		// }()

		// 无竞争，将 for 循环中i传递给 func(i int)，拷贝了一个值，没有涉及到内存地址
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
}
