package main

import (
	"fmt"
	"net"
    "os"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println("And if you try to open a file:")
	fmt.Println(os.Open("filename"))

	fmt.Println("Or access the network:")
	fmt.Println(net.Dial("tcp","google.com"))
}


/*
Welcome to the playground!
The time is 2014-03-09 18:41:08.341421 +0800 CST
And if you try to open a file:
<nil> open filename: no such file or directory
Or access the network:
<nil> missing port in address google.com

*/
