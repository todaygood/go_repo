package main

import "fmt"

func main() {
	type T struct {
		a int
		b string
	}

	t := T{27, "Margin"}
	a := []int{1, 2, 3, 4}

	var u64 uint64 = 1<<64 - 1

	fmt.Printf("%v %v %v\n", u64, t, a)

	fmt.Print(u64, " ", t, " ", a, "\n")

	fmt.Println(u64, t, a)

}
