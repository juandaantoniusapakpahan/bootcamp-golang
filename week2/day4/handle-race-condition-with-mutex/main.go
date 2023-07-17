package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mx := new(sync.Mutex)
	var x = 0
	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mx.Lock()
				x = x + 1
				mx.Unlock()
			}

		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println(x)
}
