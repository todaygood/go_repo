package main

import (
     "fmt"
	 "sync"
)

var (
	count = 5
)



func myroutine(wg *sync.WaitGroup) {

	defer (*wg).Done()
	(*wg).Add(1)
	fmt.Println("in myroutine")
}


func main() {
	var wg sync.WaitGroup

	for i:=0 ; i< count; i++ {
		go myroutine(&wg)

	}


	fmt.Println("main goroutine")

	wg.Wait()

}




