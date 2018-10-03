package main

import "fmt"
import "time"

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}

}

func main() {
	go test()

	time.Sleep(time.Second)
}
