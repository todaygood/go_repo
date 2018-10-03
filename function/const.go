package main

import "fmt"

var a,b,c int = 1,2,3

var d,e = 13,"go" //忽略类型

var _,m = "hello",58

var s string = "world" 

const PI=3.14

func main(){
    //var app int 

    //valid := false //在函数内部可以使用简短声明
/*
    fmt.Println(a)
    fmt.Println(PI)
*/
    //fmt.Printf("Hello world ,你好，世界!\n")
/*
    var c complex64 = 5+5i
    //输出: (5+5i)
    fmt.Println("Value is: %v", c)
*/
    m:=`hello
                world!`
    fmt.Printf(m)

    
}

