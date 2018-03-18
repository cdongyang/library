package chan_test

import (
	"fmt"
	"testing"
	"time"
)

// 阻塞的ch被close后会立刻可以读出类型的零值
func TestChan(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("close chan")
			}
		}
	}()
	time.Sleep(time.Second)
	close(ch)
}
