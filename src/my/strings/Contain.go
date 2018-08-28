package main

import "fmt"
import "strings"

func main() {

   s := "hello,世界!!!!!!!!"
   n := strings.Contains(s,"!")

   fmt.Println("n=",n)
   
   n = strings.Contains(s,"!!")
   fmt.Println("n=",n)

   n = strings.Contains(s,"!!&")
   fmt.Println("n=",n)
}
