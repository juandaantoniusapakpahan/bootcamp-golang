package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)
	channel := make(chan int)
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(ctx context.Context, counter <-chan int) {
		loop:
			for {
				select {
				case <-ctx.cancel():
					if err := ctx.Err(); err != nil {
						fmt.Println(err)
					}
					wg.Done()
					break loop

				case num := <-counter:
					fmt.Println("func:", num)
				}
			}
		}(ctx, channel)
		go func(t int) {
			time.Sleep(1 * time.Second)
			channel <- t
		}(i)
	}

	time.Sleep(7 * time.Second)
	cancel()
	close(channel)
	wg.Wait()
	fmt.Println("Selesai")
}
