package main

import "fmt"
import "time"

func main(){
   go fmt.Println("another goroutine")
   fmt.Println("main goroutine")

   time.Sleep(3*time.Second)

   fmt.Println(time.Second)
}



