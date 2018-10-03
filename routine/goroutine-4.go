package main

import "fmt"

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}

}

func main() {
	fmt.Println("vim-go")
	go test()
	var input string
	fmt.Scanln(&input)
}
