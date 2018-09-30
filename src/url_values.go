package main

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}

	v.Set("name", "Ava")
	v.Add("friend", "john")
	v.Add("friend", "bob")
	v.Add("friend", "Tom")

	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
}
