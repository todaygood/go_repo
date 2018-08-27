package test

// 公开函数,因为函数名的首字符是大写的。
func Even(i int) bool {
    return i % 2 == 0
}


// 私有函数
func odd (i int) bool {
return i % 2 == 1
}
