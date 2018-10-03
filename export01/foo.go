package export01

// Even is a 公开函数,因为函数名的首字符是大写的。
func Even(i int) bool {
    return i % 2 == 0
}


// odd is a 私有函数
func odd (i int) bool {
	return i % 2 == 1
}


//package name 注意不能包含"-" , "_"等特殊符号。
