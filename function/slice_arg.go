package main

import "fmt"

func modifyElem(a int) {
	a+=100
}

func modifyArray(a [5]int) {
	a=[5]int{5,6,7,8,9}
}

func main(){
   var s = [5]int{1,2,3,4,5}

   modifyElem(s[0])
   fmt.Println(s[0])

   modifyArray(s)
   fmt.Println(s)

}


/*

golang中传参都是传值，Ref https://www.jianshu.com/p/21ee9cdd2df4


*/
