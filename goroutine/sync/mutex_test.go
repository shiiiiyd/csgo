package sync_test

import (
	"sync"
	"testing"
)

func TestMutex(t *testing.T) {
	var mtx sync.Mutex
	counter := 0
	for i := 0; i < 3000; i++ {
		go func() {
			defer func() {
				mtx.Unlock()
			}()
			mtx.Lock()
			counter++
		}()
	}
}
