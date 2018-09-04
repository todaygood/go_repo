package employee

import (
	"fmt"
)


type Employee struct {
	FirstName string
	LastName string
	TotalLeaves int
	LeavesTaken int 
}


func ( e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName,e.LastName, 
				(e.TotalLeaves-e.LeavesTaken) ) 
}


// 一个结构和一个方法，就相当于一个类


