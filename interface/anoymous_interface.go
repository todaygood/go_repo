package main

import "fmt"

type Tester struct {
	s interface {
		String() string
	}
}


type User struct {
	id int 
	name string
}


func (self*User) String() string {

	return fmt.Sprintf("user %d; %s",self.id,self.name) 
}

func main(){
	t := Tester{ &User{1,"Jun"} }  // 这一句好难理解 

	fmt.Println(t.s.String()) 

}


/*

匿名接口可用作变量类型，或结构成员。
https://www.kancloud.cn/liupengjie/go/570057



*/
