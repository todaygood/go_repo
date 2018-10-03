package main 

import "fmt"
import "time"

func main(){

	t1 := time.Now()
	fmt.Println(t1)

	t2 := t1.Add(3*time.Hour) 
	fmt.Println(t2)

}

