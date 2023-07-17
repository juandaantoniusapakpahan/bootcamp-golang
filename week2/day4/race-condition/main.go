package main

import (
	"fmt"
	"time"
)

func main() {
	// race condition:

	var x = 0
	for i := 1; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				x = x + 1

				// akan ada proses penjumlahan yang nilai yang sama
				// cth: x = 40 + 1 mungkin akan ada dua atau lebih penjumlahan yang sama
			}
		}()
	}
	time.Sleep(time.Second * 4)
	fmt.Println(x) // seharusnya 1000 x 100
}
