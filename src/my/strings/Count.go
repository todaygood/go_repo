package main

import "fmt"
import "strings"

func main() {

   s := "hello,世界!!!!!!!!"
   n := strings.Count(s,"!")

   fmt.Println("n=",n)
   
   n = strings.Count(s,"!!")

   fmt.Println("n=",n)
}
