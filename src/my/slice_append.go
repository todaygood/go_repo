package main

import "fmt"

func main() {

	s0 := []int{1, 3}
	s1 := append(s0, 2)

	s2 := append(s1, 6, 10)
	s3 := append(s2, s0...)
	s4 := append(s3[3:6], s3[2:]...)

	fmt.Printf("s0=%v,s1=%v,s2=%v,s3=%v\n", s0, s1, s2, s3)
	fmt.Printf("s4=%v\n", s4)

	var t []interface{}
	t = append(t, 42, 3.1415, "foo")
	// interface{}空界面类型的数组变量，类似C语言的void＊，可以把任何类型的值放入其单元

	var b []byte
	b = append(b, "bar"...)

	fmt.Printf("t=%v,b=%v\n", t, b)

}
