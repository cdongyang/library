package chan_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 阻塞的ch被close后会立刻可以读出类型的零值
func TestChan(t *testing.T) {
	ch := make(chan struct{})
	go func() {
		for {
			<-ch
			fmt.Println("close chan")
		}
	}()
	time.Sleep(time.Second)
	close(ch)
}

func TestChanMiltiRecv(t *testing.T) {
	ch := make(chan int)
	closech := make(chan struct{})
	group := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func(idx int) {
			for {
				select {
				case elem := <-ch:
					fmt.Println("recv", idx, "get elem", elem)
					group.Done()
				case <-closech:
					fmt.Println("recv close")
					return
				}
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		group.Add(1)
		ch <- i
	}
	group.Wait()
	close(closech)
	time.Sleep(time.Second)
}

func TestChanClose(t *testing.T) {
	defer func() {
		err := recover()
		t.Log(err)
	}()
	var ch chan int
	if ch != nil {
		t.Fatal("not nil")
	}
	ch = make(chan int)
	if ch == nil {
		t.Fatal("nil")
	}
	close(ch)
	if ch != nil {
		t.Log("close chan is not nil")
	}
	close(ch)
}
