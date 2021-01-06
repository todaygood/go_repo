package main

import "fmt"

func main() {
	var a int = 10
	var x int

	fmt.Println("Please input one integer:")
	fmt.Scanln(&x)

	switch a {
	case x:
		fmt.Println("yes")

	case 103:
		fmt.Println("not yes")
	}

	fmt.Printf("input integer is:%v\n", x)
}
