package main

import "fmt"
import "regexp"
import "bytes"

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	// This finds the match for the regexp.
	fmt.Println(r.FindString("peach punch"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

}
