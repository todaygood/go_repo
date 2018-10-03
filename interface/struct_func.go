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

	fmt.Println("s.area=", s.area())
	fmt.Println("c.area=", c.area())
}

/*
Methods

Although this is better than the first version of this code, we can improve it significantly by using a special type of function known as a method:

func (c *Circle) area() float64 {
  return math.Pi * c.r*c.r
  }

  In between the keyword func and the name of the function we've added a “receiver”. The receiver is like a parameter – it has a name and a type – but by creating the function in this way it allows us to call the function using the . operator:

  fmt.Println(c.area())


注意：Method 和function的区别，中间有一个receiver

*/
