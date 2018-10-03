package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
}
