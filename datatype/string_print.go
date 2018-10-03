package main

import "fmt"


func main(){
	var greeting = "hello world"

	fmt.Printf("normal string:")
	fmt.Printf("%s\n",greeting)

    fmt.Printf("hex bytes:")
	for i:=0 ; i<len(greeting) ; i++ {
		fmt.Printf("%x ",greeting[i])
	}

	fmt.Printf("\n")



}


/*

注意：字符串文字是不可变的，因此一旦创建后，字符串文字就不能更改了。
*/
