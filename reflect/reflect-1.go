package main

import "fmt"
import "reflect"

func main() {
	var x float64 = 3.14

	fmt.Println("type of x is ", reflect.TypeOf(x))

	v := reflect.ValueOf(x)

	fmt.Println("value of x is ", v)
	fmt.Println("type of x is ", v.Type())

	fmt.Println("kind:", v.Kind())
	fmt.Println("value:", v.Float())
	fmt.Println(v.Interface())
	fmt.Printf("value is %5.2e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

}
