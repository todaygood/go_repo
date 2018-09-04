package main

import "fmt"


func MyPrint(v interface{}){
	fmt.Printf("%T:%v\n",v,v)
}

func main(){
	MyPrint(1)

	MyPrint("hello,world")

}


