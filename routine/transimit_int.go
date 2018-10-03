package main

import "fmt"


var (
	count =5 
)

func main(){
	done:= make (chan int)

	for i:=0 ; i< count ; i++ {
		go big_func(done)
	}
	
	fmt.Println("main goroutine")

	cnt:=0
	for cnt!=count {
        cnt += <-done 
	}
}


func big_func(done chan int){

	defer do_something(done) 
	fmt.Println("another goroutine")
}

func do_something(done chan int){

	done <- 1 			
}



