package main

import "fmt"

func main(){
	str := "hello" + ",world"

	fmt.Println(str)

	i :=8 

	i++

	fmt.Printf("i=%v\n",i)

	i--


	fmt.Printf("i=%v\n",i)

}

/*

//错误使用
    ++i //错误
    --i //错误
    var j int;
    j = i++ //错误
    if i++ > 10 {  //错误
        return
    }

*/

