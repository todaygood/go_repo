

# interface{} 是 任意类型 

interface{} 类型，这种类型的接口没有声明任何一个方法，俗称空接口。因为任何类型的对象都默认实现了空接口（上文提到任意类型只要实现了接口的方法就算实现了对应的接口，由于空接口没有方法，故所有类型默认都实现了空接口）在需要任意类型变量的场景下非常有用。
