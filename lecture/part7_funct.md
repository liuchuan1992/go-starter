#函数

函数支持多返回值

#结构体
* 结构体有点像java里的类
* 结构体中的方法：method  
* struct 类似class 怎么往strcut里加方法  通过显式声明reslover来实现和某个类型的组合
* 只能为同一个包中的类型定义方法
* Receiver 可以是类型的值或者指针
* 可以使用值或者指针来调用方法，编译器会自动转换
* 不存在方法重载
* 从某种意义上说，方法是函数的语法糖，因为receivcer其实就是方法所接受的第一个参数
* 方法的签名由 A & method ，即receiver和方法名所确定的  所以以下两个方法是不行的
 func (a A)method()
 func (a A)method(str string)
* 方法传递时，引用类型或者值类型的指针 操作的是底层的同一份数据  否则方法体内只是值类型的拷贝

method value & method expression


#类型别名和方法的组合
