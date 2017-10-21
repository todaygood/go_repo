package main

import "fmt"

func main() {
	i := 5

	fmt.Printf("An integer:%d, its location in memory:%p\n", i, &i)

	var iP *int
	iP = &i

	fmt.Printf("An integer:%d, its location in memory:%p,\n", *iP, iP)
}
