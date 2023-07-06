package main

import (
	"fmt"

	calculate "github.com/juandaantoniusapakpahan/countaja"
)

func main() {
	math := calculate.NewCalculate()
	fmt.Println(math.Add(1, 2, 3, 4))
	fmt.Println(math.Max(1, 2, 3, 4, 5, 2, 3, 10, 44))
	fmt.Println(math.Min(3, 4, 2, 5, 2, 4, 1, 4))
	fmt.Println(math.SortAsc(23, 2, 3, 2, 4, 5, 6345, 23, 5))
	fmt.Println(math.SortDesc(23, 2, 3, 2, 4, 5, 6345, 23, 5))

}
