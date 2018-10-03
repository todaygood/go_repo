package main

import (
	"fmt"
)

const (
	i = iota
	j = iota
	//j
	k
	l
)

func main() {
	const pi = 3.14
	const name = "margin Hu"
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
	fmt.Println("pi=", pi)
	fmt.Println("name=", name)

}
