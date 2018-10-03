package main

import "fmt"

type square struct{ r int }
type circle struct{ r int }

func (s square) area() int { return s.r * s.r }
func (c circle) area() int { return c.r * 3 }

func main() {
	s := square{1}
	c := circle{2}
	a := [2]interface{}{s, c}
	fmt.Println(s, c, a)

	sum := 0
	for _, t := range a {
		switch v := t.(type) {
		case square:
			sum += v.area()
			fmt.Println("square area=",sum)
		case circle:
			sum += v.area()
			fmt.Println("circle area=",sum)
		}
	}
	fmt.Println(sum)
}

/*
For instance, all types implement the empty interface:

interface{}

Ref: https://golang.org/ref/spec#Interface_types

*/
