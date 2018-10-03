package employee

import (
	"fmt"
)


type employee struct {
	firstName string
	lastName string
	totalLeaves int
	leavesTaken int 
}

func  New(firstName string,lastName string,totalLeaves int,leavesTaken int) employee {
	e := employee{firstName,lastName,totalLeaves,leavesTaken}
	return e 
}

func ( e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.firstName,e.lastName, 
				(e.totalLeaves-e.leavesTaken) ) 
}


// 一个结构和一个方法，就相当于一个类, 用小写来实现不让外部访问的目的。


