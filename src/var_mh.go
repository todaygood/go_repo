
package main


import (
    "fmt"
    "runtime"
    "os"
    )

func main(){

    var a int 

    var b bool 

    var str string

    a = 15 
    b = true 
    str = "hello,world!"

    fmt.Printf("a=%d, b=%d,str=%s\n",a,b,str)

    var goos string = runtime.GOOS

    fmt.Printf("The operating system is %s \n", goos ) 

    path := os.Getenv("PATH") 

    fmt.Printf("Path is %s \n",path) 

}
