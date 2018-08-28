package main

import "fmt"
import "errors"

type CustomError struct {
	infoa string
	infob string
	Err   error
}

func (cerr CustomError) Error() string {
	errorInfo := fmt.Sprintf("infoa:%s,infob:%s, origin err info:%s\n",
		cerr.infoa, cerr.infob, cerr.Err.Error())
	return errorInfo

}

func main() {

	var err error = errors.New("this is a new error")
	fmt.Println(err.Error())

	err = fmt.Errorf("%s", "the error test for fmt.Errorf")
	fmt.Println(err.Error())

	err = &CustomError{
		infoa: "error info a",
		infob: "error info b",
		Err:   errors.New("test custom err"),
	}

	fmt.Println(err.Error())

}
