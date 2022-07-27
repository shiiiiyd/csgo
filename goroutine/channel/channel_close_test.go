// 生产的消息的chan关闭后，在发送消息出现panic，
// 使用 close()函数关闭chan通道
package channel_test

import (
	"fmt"
	"sync"
	"testing"
)

// 生产
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// ch 中消息发送完毕后关闭，如果再次向该chan中发送消息会出现panic
		close(ch)
		wg.Done()
	}()
}

// 消费
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func dataReceiverNum(num int, ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Printf("receiver(%d)>> %d\n", num, data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestChanClose(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

func TestChanCloseNum(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)

	wg.Add(1)
	dataReceiverNum(1, ch, &wg)
	wg.Add(1)
	dataReceiverNum(2, ch, &wg)
	wg.Wait()
}
