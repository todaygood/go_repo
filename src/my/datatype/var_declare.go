package main

import "fmt"

var x, y int 

var (
	a int 
	b bool
)   //这种写法一般用于声明全局变量


var c,d int = 1, 2 

var e,f = 123, "hello"  //自动推断出类型


func main(){
	g,h := 789, "world"  //简略声明只能用于函数体内

	fmt.Println(x,y,a,b,c,d,e,f,g,h)


}



/*
Ref  http://www.runoob.com/go/go-variables.html

*/
