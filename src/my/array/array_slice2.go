package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

        fmt.Println("print numbers:")
	printSlice(numbers)

	fmt.Println("numbers=", numbers)

	fmt.Println("numbers[1:4]=", numbers[1:4])

	fmt.Println("numbers[:4]=", numbers[:4])

	fmt.Println("numbers[4:]=", numbers[4:])

	var numbers1 []int
	numbers1 = make([]int, 0, 5)

	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2) // why cap== 10 ? 

	numbers3 := numbers[3:5]
	printSlice(numbers3)  // why cap== 7 ? 
}

func printSlice(x []int) {
	fmt.Printf("len=%v,cap=%v,slice=%v\n", len(x), cap(x), x)
}



/*

cap 是Len - minIndex
len(x) 返回切片中的元素数量
cap(x) 返回切片的容量大小。

func make(t Type, size ...IntegerType) Type
说明：...表示size参数是一个不定参数

:= 用于变量的简短声明，参见https://blog.csdn.net/wfyeshi/article/details/60572739
操作 := 只能用于方法内部, 声明并初始化新的变量,不能用于已声明变量赋值
:= 操作左边必须要有新变量, 那么多个变量初始化只需要满足左边至少有一个新变量即可

该例子说明在golang中，函数可以先被调用，后被定义。

*/
