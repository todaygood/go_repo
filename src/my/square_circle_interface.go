package main

import "fmt"

type Rect struct {
	r int
}

type Circle struct {
	r int
}

func (s Rect) area() int {
	return s.r * s.r
}

func (c Circle) area() int {
	return c.r * 3
}

func main() {
	s := Rect{1}
	c := Circle{3}

	fmt.Println(s, c, s.area()+c.area())
}
