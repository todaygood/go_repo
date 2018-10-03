package main


import "fmt"
import "github.com/todaygood/go_repo/export01"

func main(){

   i := 5 
    

   fmt.Printf("Is %d even ? %v \n",i, export01.Even(i))
   //fmt.Printf("Is %d even ? %v \n",i, export01.odd(i))

}

