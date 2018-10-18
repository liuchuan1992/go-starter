#interface接口
接口配合结构一起使用

Comma-ok断言的语法是：value, ok := element.(T)。
	element必须是接口类型的变量，T是普通类型。
	如果断言失败，ok为false，否则ok为true并且value为变量的值。