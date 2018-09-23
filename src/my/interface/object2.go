package main

import "fmt"


func PrintAll(vals [] interface{}) {
	for _,val := range vals {
		fmt.Println(val)
	}
}

func main(){
	names := []string{"tom","david","oscar"}
	vals := make([]interface{},len(names))
	for i,v := range names {
		vals[i]=v 
	}
	PrintAll(vals)
}

/*
    利用 interface{}进行类型转换
*/

