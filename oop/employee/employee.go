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



/*

一个结构和一个方法，就相当于一个类

  struct的方法,看起来仅仅是放置位置的不同，其实是设计理念的不同。将方法放在类定义里面，意味着方法是类不可分割的一部分，类的最小单位就是数据成员和当前定义的所有操作，你要认识这个类，必须一次性认识这个类中定义的所有的东西。相反，先定义类的数据结构，然后像搭积木一样将目前需要的方法一个一个地进行绑定，你便可以根据需求对类进行扩展。传统的类定义是你必须一开始便想好这个类有哪些操作，一旦类定义好了，类就成了你定义的样子，再无其他可能。golang的这种开放式扩展定义方式，使得类更加具有生命力，你不必一开始就设计好一切(往往也很难做到)，类会随着你的实现思路逐渐成长为你想要的那个样子。

*/
