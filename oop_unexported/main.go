package main 

import "oop_unexported/employee"

func main() {
	testcase1()
	testcase2()
}


func testcase2() {
	e := employee.New( "", "", 30, 10 )
	e.LeavesRemaining()
}

func testcase1() {
	e := employee.New( "sam", "Smith", 30, 20 )
	e.LeavesRemaining()
}




