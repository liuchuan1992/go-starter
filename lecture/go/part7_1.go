package main

import "fmt"

type A struct {
	name string
}

type C struct{
	name string
}

type TZ int

func main()  {
	a := A{"lvcf a "}
	a.sayHello("你好呀")
	println(a.name)

	c := C{"lvcf c"}
	//method value
	c.sayHello("你好呀")
	//method expression
	(C).sayHello(c,"你好呀")
	println(c.name)

	var tz TZ
	tz.Print()
}

/**
	定义struct内的方法
 */
func (a *A)sayHello(str string)  {
	a.name = "liuchuan A"
	fmt.Println("hello,",a.name,str)
}

func (c C)sayHello(str string)  {
	c.name = "liuchuan C"
	fmt.Println("hello,",c.name,str)
}

func (tz TZ)Print(){
	fmt.Println("TZ")
}