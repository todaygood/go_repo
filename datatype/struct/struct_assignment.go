package main

import "fmt"

type person struct {
	name string 
	age int 
}


func main(){
	s := person{name: "Margin",age: 50}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp)


}

