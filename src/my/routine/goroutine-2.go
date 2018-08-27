package main

import "fmt"

func test(done chan bool) {
/*
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}
*/
	done <- true

}

func main() {
        var result bool

	fmt.Println("in main")

	done := make(chan bool)
	go test(done)

	result = false
	result = <-done
	fmt.Println("result=",result)
}
