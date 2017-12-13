package main

import "fmt"

func main() {

	var countryCaptitalMap map[string]string

	countryCaptitalMap = make(map[string]string)
	countryCaptitalMap["France"] = "Paris"
	countryCaptitalMap["Italy"] = "Rome"
	countryCaptitalMap["Japan"] = "Tokyo"

	for country := range countryCaptitalMap {
		fmt.Println("Captital of", country, "is:", countryCaptitalMap[country])

	}
}
