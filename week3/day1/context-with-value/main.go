package main

import (
	"context"
	"fmt"
)

func main() {

	root := context.Background()
	ctx1 := context.WithValue(root, "ctx1", "Contex 1")
	ctx2 := context.WithValue(root, "ctx2", "Contex 2")
	ctx11 := context.WithValue(ctx1, "ctx11", "Context-1.1")
	ctx12 := context.WithValue(ctx1, "ctx12", "Context-1.2")

	fmt.Println(root)
	fmt.Println(ctx1)
	fmt.Println(ctx2)
	fmt.Println(ctx11)
	fmt.Println(ctx12)
	fmt.Println()
	fmt.Println(root.Value("ctx1"))

}
