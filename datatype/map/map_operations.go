package main

import "fmt"

func main(){
	
	amap := make(map[string]int)

	amap["hello"]=3
	amap["world"]=5


	fmt.Println("map is",amap)
	fmt.Println("len is ",len(amap))

	for i,v := range amap {
		if v > 4 {
			fmt.Printf("found,i=%v, v=%v\n",i,v)
		}
	}

	delete(amap,"hello")


	fmt.Println("map is",amap)
	fmt.Println("len is ",len(amap))


	n := map[string]int{"a":1,"b":3}
	fmt.Println("map n is: ", n)

}
