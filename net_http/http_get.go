package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://www.baidu.com"
	httpGet(url)
}

func httpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Fatalln("Error http.Get")
	}
	fmt.Println()
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		log.Fatalln("Error ReadAll")
	}

	fmt.Println(string(body))
}
