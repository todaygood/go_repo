package main

import "fmt"

func Greeting(prefix string, who ...string) {
	for _, person := range who {
		fmt.Println(prefix, person)

	}
}

func main() {
	s := []string{"James", "Jack"}

	Greeting("hello,", s...)

}
