package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

func main() {
	timer := time.Now().Unix()
	if timer%Female == 0 {
		fmt.Println("%d is Female", timer)
		fmt.Printf("%d is Female\n", timer)
	} else {
		fmt.Println("%d is Man", timer)
		fmt.Printf("%d is Man\n", timer)
	}
}
