package goroutine

import (
	"fmt"
	"testing"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
