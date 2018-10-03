package main

import "fmt"


type Student struct {
	sex string;
    id  int; 
}

func main() {
	i := 5

	fmt.Printf("An integer:%d, its location in memory:%p\n", i, &i)

	var iP *int
	iP = &i

	fmt.Printf("An integer:%d, its location in memory:%p,\n", *iP, iP)


	var Tom Student= Student{"male",1}

	fmt.Println("Tom is" , Tom.sex,(*(&Tom)).id)
}



/*
golang的指针不使用符号 "->", 只有"&","*"两个符号可以用。

ref: http://www.runoob.com/go/go-pointers.html
*/


