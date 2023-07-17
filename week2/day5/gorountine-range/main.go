package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	name := GetAll()
	data := make(chan any, 2)

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go GetName(name, data, wg)
	go GetAge(name, data, wg)

	wg.Wait()
	close(data)
	for v := range data {
		fmt.Println("STEP:", v)
	}

	fmt.Println(time.Since(start))

}
func GetName(name string, data chan any, w *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	data <- "GGWP"
	w.Done()
}

func GetAge(age string, data chan any, w *sync.WaitGroup) {
	time.Sleep(time.Microsecond * 100)
	data <- "11"
	w.Done()
}

func GetAll() string {
	return "TEST"
}
