package main

import (
	"fmt"
	"strings"
)

func Voko(str string) {
	input1 := strings.ToLower(str)
	vocal := "aiueo"

	if strings.Contains(vocal, input1) {
		fmt.Println(input1, "merupakan huruf vocals")
	} else {
		fmt.Println(input1, "merupakan huruf konsonan")
	}
}

func main() {

	fmt.Println("Enter Character")
	var input string

	fmt.Scanln(&input)
	Voko(input)
}
