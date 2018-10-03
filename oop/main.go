package main 

import "oop/employee"

func main() {
	testcase1()
	testcase2()
}

func testcase2() {
	var e employee.Employee 

	e.LeavesRemaining()
}


func testcase1() {
	e := employee.Employee { 
		FirstName: "sam",
		LastName: "Smith",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	e.LeavesRemaining()

}




