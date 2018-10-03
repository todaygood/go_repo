package main

import "fmt"

func mytest(args ...string) {
	for _, v := range args {
		fmt.Println(v)
	}

}

func main() {
	var stress = []string{
		"qwr",
		"234",
		"yui",
	}

	mytest(stress...)
}
