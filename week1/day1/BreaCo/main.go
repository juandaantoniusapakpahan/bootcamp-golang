package main

import "fmt"

func BreaCo() []int {
	count := 0
	i := 1
	var result []int
	for {
		if i%3 == 0 || i%5 == 0 {
			result = append(result, i)
			count++
		}
		i = i + 2

		if count == 5 {
			break
		} else {
			continue
		}
	}
	return result
}

func main() {
	result := BreaCo()
	fmt.Println(result)
}
