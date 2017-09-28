package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	const pi float64 = 3.14
	var a int

	var b bool

	var str string

	a = 15
	b = true
	str = "hello,world!"

	fmt.Printf("a=%d, b=%d,str=%s,pi=%f\n", a, b, str, pi)

	var goos string = runtime.GOOS

	fmt.Printf("The operating system is %s \n", goos)

	path := os.Getenv("PATH")

	fmt.Printf("Path is %s \n", path)

}
