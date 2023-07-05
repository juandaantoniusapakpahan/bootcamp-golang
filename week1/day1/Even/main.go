package main

import "fmt"

func Even(a, b int) []int {
	var result []int
	for a <= b {
		if a%2 == 0 {
			result = append(result, a)
		}
		a++
	}
	return result
}

func main() {
	even := Even(10, 100)
	fmt.Println(even)
}
