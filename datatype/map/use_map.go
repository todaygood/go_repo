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
 

    //查看元素是否在map中存在
	capital,ok := countryCaptitalMap["Fance"]
    fmt.Println(capital,ok)
}

/*
map是一种无序的键值对的集合。map最重要的一点是通过key来快速检索数据。

*/
