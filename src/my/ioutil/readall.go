package main

import (
    "fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main(){
	r := strings.NewReader("Go language")

	b,err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(b)
	fmt.Printf("%s\n",b)

}

/*
log.Fatal  is equivalent to Print() followed by a call to os.Exit(1).

*/
