package main

import "fmt"

func main() {

	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	var b = make([]byte, 5)

	n1 := copy(s, a[0:])
	fmt.Printf("s=%v,n1=%v\n", s, n1) //n1=6 , s={0,1,2,3,4,5}

	n2 := copy(s, s[2:])
	fmt.Printf("s=%v,n2=%v\n", s, n2) //n2=4, s={2,3,4,5,4,5}

	n3 := copy(b, "hello,world!") //b=Ascii("hello,")  n3=5
	fmt.Printf("b=%v,n3=%v\n", b, n3)

}
