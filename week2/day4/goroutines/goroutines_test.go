package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	mx := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	channel := make(chan int, 10)
	go func() {
		for i := 1; i <= 100; i++ {
			wg.Add(1)
			mx.Lock()
			channel <- i
			mx.Unlock()
			wg.Done()
		}
	}()

	for i := 1; i <= 100; i++ {
		ex := <-channel
		fmt.Println(ex)
	}
	time.Sleep(1 * time.Second)
}

func TestBuffer(t *testing.T) {
	channel := make(chan string, 3) // Panjang/kapasita channel yaitu 3,
	channel <- "babi"
	channel <- "anjing"
	channel <- "cicak"

	// Jika data disimpan lebih dari 3 maka akan error/tidak jalan

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	fmt.Println("GGWP")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			channel <- i
		}
		close(channel)
	}()

	for v := range channel {
		fmt.Println(v)
	}
	fmt.Println("Selesai")
}
