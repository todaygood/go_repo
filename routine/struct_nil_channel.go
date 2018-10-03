package main

import "fmt"

func main(){

	done := make(chan struct {})

	go func(){

		defer close(done)
		fmt.Println("another goroutine")
        }() //go语句支持匿名函数，所以需要加()表示在调用一个函数

	fmt.Println("main goroutine")

	<- done 
}


/*
https://zhengheng.me/2016/10/24/goroutine-notes/

channel需要使用make函数生成。 

<- done 会在main函数中阻塞， 直到在goroutine中关闭channel
*/
