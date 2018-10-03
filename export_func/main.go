package main


import "fmt"
import "github.com/todaygood/go_repo/export_func"

func main(){

   i := 5 
    

   fmt.Printf("Is %d even ? %v \n",i, export_func.Even(i))
   //fmt.Printf("Is %d even ? %v \n",i, export_func.odd(i))

}

