package main

import "fmt"

func main() {
	var numbers []int
	printSlice(numbers)
	if numbers == nil {
		fmt.Println("slice is nil")
	}
	fmt.Println("vim-go")
}

func printSlice(x []int) {
	fmt.Printf("len=%v,cap=%v,slice=%v\n", len(x), cap(x), x)

}
