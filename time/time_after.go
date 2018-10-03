package main 

import "fmt"
import "time"

func main(){
	fmt.Println("the 1")
	fmt.Println(time.Now())
	tc := time.After(3*time.Second) 

	fmt.Println("the 2")
	fmt.Println("the 3")

	t:=<-tc 

	fmt.Println("the 4")
	fmt.Println(t)

}

