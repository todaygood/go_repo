package main

import "fmt"

func main() {
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型是: %T", i)

	case int:
		fmt.Printf("x is int")
	case float64:
		fmt.Printf("x is float64")
	default:
		fmt.Printf("unknown type")

	}

	fmt.Println("vim-go")
}
