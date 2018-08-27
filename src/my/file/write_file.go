package main

import "fmt"
import "os"

/*
go中的defer用法：表示延迟做一些事情，参见https://xiaozhou.net/something-about-defer-2014-05-25.html
常用于close file handle,等类似析构函数

同一个函数中的多个defer，是先进后出的。
当defer被声明时，其参数就会被实时解析

*/

func main() {
	f := createFile("defer-test.txt")
	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
